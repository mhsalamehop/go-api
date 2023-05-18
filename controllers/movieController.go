package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/mhsalamehop/go-api/models"
	"github.com/mhsalamehop/go-api/store"
	"github.com/mhsalamehop/go-api/utils"
)

type MovieResult struct {
	Page         int64           `json:"page"`
	Results      []models.Movies `json:"results"`
	TotalPages   int64           `json:"total_pages"`
	TotalResults int64           `json:"total_results"`
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	db:=store.OpenConnection()
	// sli := []string{"id","overview","title"}
	valueStrings := []string{}
	valueArgs := []interface{}{}
	txn,err := db.Begin()
	if err != nil {
		utils.LogInfo(err.Error())
	}

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
	c:=0
	for _, w := range  movieResult.Results{
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", c+1, c+2, c+3))    
		valueArgs = append(valueArgs, w.Id)
		valueArgs = append(valueArgs, w.Title)
		valueArgs = append(valueArgs, w.Overview)
		c+=3
	  }
	smt := `INSERT INTO api_movies (id, title, overview) VALUES %s ON CONFLICT (id) DO UPDATE SET overview = excluded.overview`
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ", "))
	_,err = txn.Exec(smt, valueArgs...)
	if err != nil {
		utils.LogInfo(err.Error())
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}

	data, err := json.Marshal(movieResult.Results)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)

}


