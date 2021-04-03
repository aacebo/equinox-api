package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/aacebo/equinox-api/src/api"
	"github.com/aacebo/equinox-api/src/db"
	"github.com/aacebo/equinox-api/src/logger"
)

func main() {
	var env = os.Getenv("GO_ENV")
	var err = godotenv.Load(fmt.Sprintf(".env.%s", env))
	var log = logger.New("main")

	if err != nil {
		log.Error(err)
	}

	var port = os.Getenv("PORT")

	log.Infof("running on %s", env)

	var db, dberr = db.Connect()

	if dberr != nil {
		log.Error(dberr)
	}

	log.Info("connected to database...")

	defer db.Close()

	var router, routererr = api.Router(env, db)

	if routererr != nil {
		log.Error(routererr)
	}

	router.Run(fmt.Sprintf("localhost:%s", port))
}
