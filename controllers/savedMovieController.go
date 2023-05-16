package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhsalamehop/go-api/models"
	"github.com/mhsalamehop/go-api/store"
)

func GetSavedMovies(w http.ResponseWriter, r *http.Request) {
	var movie models.SavedMovies
	var savedInfoJson []byte
	db := store.OpenConnection()
	querystr := `SELECT id, title, overview, saved_info FROM saved_movies`
	rows, err := db.Query(querystr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var movies []models.SavedMovies
	for rows.Next() {
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Overview, &savedInfoJson)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(savedInfoJson,&movie.SavedInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movies = append(movies, movie)
	}
	data, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.SavedMovies
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	saveInfoJson , err := json.Marshal(movie.SavedInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := store.OpenConnection()
	queryStr := `INSERT INTO saved_movies (title, overview, saved_info) VALUES($1, $2, $3) RETURNING id`
	err = db.QueryRow(queryStr, movie.Title, movie.Overview, saveInfoJson).Scan(&movie.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.SavedMovies
	id := mux.Vars(r)["id"]
	movie.Id = id
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.Unmarshal(body, &movie); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db := store.OpenConnection()
	queryStr := `UPDATE saved_movies SET title=$1, overview=$2 WHERE id=$3`
	db.QueryRow(queryStr, movie.Title, movie.Overview, id)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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
