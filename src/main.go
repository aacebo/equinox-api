package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/aacebo/equinox-api/src/api"
	"github.com/aacebo/equinox-api/src/db"
)

func main() {
	var env = os.Getenv("GO_ENV")
	var err = godotenv.Load(fmt.Sprintf(".env.%s", env))

	if err != nil {
		log.Fatal(err)
	}

	var port = os.Getenv("PORT")

	log.Infof("running on %s", env)

	var db = db.Conn()

	defer db.Close()

	var router = api.Router(env)

	router.Run(fmt.Sprintf("localhost:%s", port))
}
