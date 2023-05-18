package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/mhsalamehop/go-api/models"
	"github.com/mhsalamehop/go-api/utils"
)

type MovieResult struct {
	Page         int64           `json:"page"`
	Results      []models.Movies `json:"results"`
	TotalPages   int64           `json:"total_pages"`
	TotalResults int64           `json:"total_results"`
}

func GetMoviesFromAPI() (MovieResult, error) {
	err := godotenv.Load(".env")
	if err != nil {
		utils.LogError("loading .env file with err %s", err.Error())
	}
	url := "https://api.themoviedb.org/3/movie/popular?language=en-US&page=1"
	token := os.Getenv("MYMOVIEDB_TOKEN")
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	var movieResult MovieResult
	if err != nil {
		utils.LogError(err.Error())
		return movieResult, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		utils.LogError(err.Error())
		return movieResult, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		utils.LogError(err.Error())
		return movieResult, err
	}
	err = json.Unmarshal(body, &movieResult)
	if err != nil {
		utils.LogError(err.Error())
		return movieResult, err
	}
	return movieResult, nil

}

func BackFill(movieResult MovieResult, db *sql.DB) error {
	valueStrings := []string{}
	valueArgs := []interface{}{}
	txn, err := db.Begin()
	if err != nil {
		utils.LogError(err.Error())
		return err
	}
	c := 0
	for _, w := range movieResult.Results {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", c+1, c+2, c+3))
		valueArgs = append(valueArgs, w.Id)
		valueArgs = append(valueArgs, w.Title)
		valueArgs = append(valueArgs, w.Overview)
		c += 3
	}
	smt := `INSERT INTO api_movies (id, title, overview) VALUES %s ON CONFLICT (id) DO UPDATE SET overview = excluded.overview`
	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ", "))
	_, err = txn.Exec(smt, valueArgs...)
	if err != nil {
		utils.LogError(err.Error())
		return err
	}
	err = txn.Commit()
	if err != nil {
		utils.LogError(err.Error())
		return err
	}
	return nil
}
