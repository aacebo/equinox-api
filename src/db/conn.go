package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

var DB *sql.DB

func Conn() *sql.DB {
	if DB == nil {
		var uri = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("PG_HOST"),
			5432,
			os.Getenv("PG_USER"),
			os.Getenv("PG_PASSWORD"),
			os.Getenv("PG_DATABASE"),
		)

		var db, err = sql.Open("postgres", uri)

		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()

		if err != nil {
			log.Fatal(err)
		}

		DB = db
		log.Infoln("connected to database...")
	}

	return DB
}
