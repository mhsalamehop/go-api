package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhsalamehop/go-api/controllers"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/users", controllers.GetUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{user_id}", controllers.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/signup", controllers.Signup).Methods(http.MethodPost)
	r.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
}
