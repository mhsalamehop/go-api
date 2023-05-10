package controllers

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
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
	fmt.Println("I am here")
	db := store.OpenConnection()
	var user models.UserModel
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password),14)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	} 
	fmt.Printf("%v -> %v\n",user.Password,string(hashedPassword))
	querystr := `INSERT INTO users_table (email, username, password) VALUES ($1, $2, $3) RETURNING id`

	rows,err := db.Query(querystr,user.Email,user.Username,user.Password)
	if err != nil {
		http.Error(w,"signup error" + err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()
}

func Login(w http.ResponseWriter, r *http.Request) {

}
