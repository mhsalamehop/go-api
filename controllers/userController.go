package controllers

import (
	// "encoding/json"
	"encoding/json"
	"fmt"

	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhsalamehop/go-api/models"
	"github.com/mhsalamehop/go-api/store"
	"github.com/mhsalamehop/go-api/utils"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	utils.LogInfo("Getting users")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Println(params)
	utils.LogInfo("Getting user data with id %v", params["user_id"])
}

func Signup(w http.ResponseWriter, r *http.Request) {
	db := store.OpenConnection()
	var user models.UserModel
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// json.NewDecoder(r.Body).Decode(&user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	user.Password = string(hashedPassword)
	querystr := `INSERT INTO users_table (email, username, password, role) VALUES ($1, $2, $3, $4)`

	db.QueryRow(querystr, user.Email, user.Username, hashedPassword, user.Role)
	if err != nil {
		http.Error(w, "signup error "+err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := store.OpenConnection()
	// var loginInfo models.LoginInfo
	var user models.UserModel
	username, password, ok := r.BasicAuth()
	// body, err := ioutil.ReadAll(r.Body)

	// if err := json.Unmarshal(body,&loginInfo); err != nil {

	// }
	if !ok {
		http.Error(w, "no auth header", http.StatusForbidden)
		return
	}
	if username == "" || password == "" {
		http.Error(w, "please provide username and password", http.StatusForbidden)
		return
	}
	querystr := `SELECT username, password, email, id, role FROM users_table WHERE email=$1`
	rows, err := db.Query(querystr, username)
	if err != nil {
		http.Error(w, "sql here "+err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&user.Username, &user.Password, &user.Email, &user.Id, &user.Role); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	token, err := utils.GetToken(user.Email, user.Role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	updateQuery := `UPDATE users_table SET token=$1 WHERE id=$2 RETURNING token`
	db.QueryRow(updateQuery, token, user.Id).Scan(&user.Token)

	json.NewEncoder(w).Encode(user)
}
