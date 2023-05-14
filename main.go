package main

import (
	// "fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	// "github.com/mhsalamehop/go-api/middleware"
	"github.com/mhsalamehop/go-api/routes"
	// "github.com/mhsalamehop/go-api/store"
	"github.com/mhsalamehop/go-api/utils"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		utils.LogError("loading .env file with err %s", err.Error())
	}
	port := os.Getenv("PORT")

	r := mux.NewRouter()
	routes.UserRoutes(r)
	routes.SavedMovieRouter(r)
	routes.MoviesRoutes(r)
	// r.Use(middleware.Authentication())

	// routes.MoiveRoutes(r)
	// routes.SavedMovieRoutes(r)
	utils.LogInfo("running server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
