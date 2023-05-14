package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhsalamehop/go-api/models"
	"github.com/mhsalamehop/go-api/store"
)

func GetSavedMovies(w http.ResponseWriter, r *http.Request) {
	var movie models.SavedMovies
	db := store.OpenConnection()
	querystr := `SELECT id, title, overview FROM saved_movies`
	rows, err := db.Query(querystr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()
	var movies []models.SavedMovies
	for rows.Next() {
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Overview)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		movies = append(movies, movie)
	}
	data, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("role")
	fmt.Println(role)
	if role != "admin" {
		http.Error(w, "Admins only can add movies", http.StatusUnauthorized)
	}
	var movie models.SavedMovies
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := json.Unmarshal(body, &movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	db := store.OpenConnection()
	queryStr := `INSERT INTO saved_movies (title, Overview) VALUES($1, $2) RETURNING id`
	db.QueryRow(queryStr, movie.Title, movie.Overview).Scan(&movie.Id)
	data, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	role := r.Header.Get("role")
	var movie models.SavedMovies
	id := mux.Vars(r)["id"]
	movie.Id = id
	if role != "admin" {
		http.Error(w, "Admins only can add movies", http.StatusUnauthorized)
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := json.Unmarshal(body, &movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	db := store.OpenConnection()
	queryStr := `UPDATE saved_movies SET title=$1, overview=$2 WHERE id=$3`
	db.QueryRow(queryStr, movie.Title, movie.Overview, id)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	data, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	db := store.OpenConnection()
	queryStr := `DELETE FROM saved_movies WHERE id=$1`
	db.QueryRow(queryStr,id)
	w.Write([]byte("Succefulli deleted movie with id="+id))
}
