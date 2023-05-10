package store

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/lib/pq"
	"github.com/mhsalamehop/go-api/utils"
)

var (
	once sync.Once
	db   *sql.DB
)

func OpenConnection() *sql.DB {
	if db == nil {
		once.Do(func() {
			var err error
			db, err = sql.Open("postgres", os.Getenv("POSTGRES_URI"))
			if err != nil {
				utils.LogError("failed to connect %v", err)
				return
			}
			connectionError := db.Ping()
			if connectionError != nil {
				utils.LogError("%v", connectionError)
			}
		})
	}
	return db
}
