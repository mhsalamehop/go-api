package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhsalamehop/go-api/controllers"
)

func MoviesRoutes(r *mux.Router){
	r.HandleFunc("/movies",controllers.GetMovies).Methods(http.MethodGet)
}