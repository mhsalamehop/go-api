package main

import (
	"github.com/mhsalamehop/go-api/store"
	"github.com/mhsalamehop/go-api/utils"
)

func main() {
	movieResult, err := GetMoviesFromAPI(10)
	if err != nil {
		utils.LogInfo(err.Error())
	}
	db := store.OpenConnection()
	err = BackFill(*movieResult, db)
	if err != nil {
		utils.LogError(err.Error())
	}
}
