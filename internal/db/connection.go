package db

import (
	"database/sql"
	"hoc-gin/internal/config"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	connStr := config.NewConfig().DNS()

	var err error
	DB, err = sql.Open("postgres", connStr) // go get github.com/lib/pq
	if err != nil {

		log.Fatal("unable to use data source name", err)
	}
	defer DB.Close()

	return nil
}
