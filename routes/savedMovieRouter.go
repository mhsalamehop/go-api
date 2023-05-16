package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhsalamehop/go-api/controllers"
	"github.com/mhsalamehop/go-api/middleware"
)


func SavedMovieRouter(r *mux.Router){
	var amw middleware.Authorization
	r.HandleFunc("/fav",controllers.GetSavedMovies).Methods(http.MethodGet)
	r.Handle("/fav",amw.IsAuthorized(http.HandlerFunc(controllers.AddMovie))).Methods(http.MethodPost)
	r.Handle("/fav/{id}",amw.IsAuthorized(http.HandlerFunc(controllers.UpdateMovie))).Methods(http.MethodPut)
	r.Handle("/fav/{id}",amw.IsAuthorized(http.HandlerFunc(controllers.DeleteMovie))).Methods(http.MethodDelete)
}