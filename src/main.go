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

	var db = db.Connect()

	log.Info("connected to database...")

	defer db.Close()

	var router = api.Router(env, db)

	router.Run(fmt.Sprintf("localhost:%s", port))
}
