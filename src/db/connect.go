package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/aacebo/equinox-api/src/logger"
)

var log = logger.New("db")

func Connect() *sql.DB {
	var uri = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("PG_HOST"),
		5432,
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DATABASE"),
	)

	var db, err = sql.Open("postgres", uri)

	if err != nil {
		log.Error(err)
	}

	err = db.Ping()

	if err != nil {
		log.Error(err)
	}

	return db
}
