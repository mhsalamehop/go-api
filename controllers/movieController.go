package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/mhsalamehop/go-api/models"
)

type MovieResult struct {
	Page         int64           `json:"page"`
	Results      []models.Movies `json:"results"`
	TotalPages   int64           `json:"total_pages"`
	TotalResults int64           `json:"total_results"`
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	// var movie models.Movies
	url := "https://api.themoviedb.org/3/movie/popular?language=en-US&page=1"
	token := os.Getenv("MYMOVIEDB_TOKEN")
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	var movieResult MovieResult
	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &movieResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(movieResult.Results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)

}
