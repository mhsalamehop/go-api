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

func GetMoviesFromAPI(pages int) (*[]MovieResult, error) {
	err := godotenv.Load(".env")
	if err != nil {
		utils.LogError("loading .env file with err %s", err.Error())
	}
	var movies []MovieResult
	for i := 1 ; i <= pages ; i++ {

		url := fmt.Sprintf("https://api.themoviedb.org/3/movie/popular?language=en-US&page=%d",i)
		token := os.Getenv("MYMOVIEDB_TOKEN")
		client := http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		var movieResult MovieResult
		if err != nil {
			utils.LogError(err.Error())
			return nil, err
		}
		req.Header.Add("accept", "application/json")
		req.Header.Add("Authorization", "Bearer "+token)
	
		res, err := client.Do(req)
		if err != nil {
			utils.LogError(err.Error())
			return nil, err
		}
		defer res.Body.Close()
	
		body, err := io.ReadAll(res.Body)
		if err != nil {
			utils.LogError(err.Error())
			return nil, err
		}
		err = json.Unmarshal(body, &movieResult)
		if err != nil {
			utils.LogError(err.Error())
			return nil, err
		}
		movies = append(movies, movieResult)
	}
	return &movies, nil

}

func BackFill(movieResult []MovieResult, db *sql.DB) error {
	valueStrings := make([]string,0,len(movieResult))
	valueArgs := make([]interface{},0, 3*len(movieResult))
	c := 0
	for i := 0 ; i < len(movieResult) ; i++ {
		var x = movieResult[i].Results
		for _, val := range x {
			str := fmt.Sprintf("($%d, $%d, $%d)", c*3+1,  c*3+2, c*3+3)
			
			valueStrings = append(valueStrings, str)
			valueArgs = append(valueArgs, val.Id)
			valueArgs = append(valueArgs, val.Title)
			valueArgs = append(valueArgs, val.Overview)
			c++
			if len(valueArgs) >= 65000  || i == len(movieResult)-1{
				smt := `INSERT INTO api_movies (id, title, overview) VALUES %s ON CONFLICT (id) DO UPDATE SET overview = excluded.overview, title = excluded.title`
				smt = fmt.Sprintf(smt, strings.Join(valueStrings, ", "))
				_, err := db.Query(smt, valueArgs...)
				if err != nil {
					utils.LogError(err.Error())
					return err
				}
				valueStrings = make([]string,0,len(movieResult))
				valueArgs = make([]interface{},0, 3*len(movieResult))
				c = 0
			}
		}
	}

	return nil
}
