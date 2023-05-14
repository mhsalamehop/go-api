package controllers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/mhsalamehop/go-api/models"
)

type MovieResult struct {
	Results []models.Movies `json:"results"`
}

func GetMovies(w http.ResponseWriter, r *http.Request){
	// var movie models.Movies
	url := "https://api.themoviedb.org/3/movie/popular?language=en-US&page=1"
	token := os.Getenv("MYMOVIEDB_TOKEN")
	req, err := http.NewRequest("GET",url,nil)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	req.Header.Add("accept","application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
	}
	defer res.Body.Close()

	var movieResult MovieResult
	err = json.NewDecoder(res.Body).Decode(&movieResult)
	if err !=nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(movieResult.Results)
}