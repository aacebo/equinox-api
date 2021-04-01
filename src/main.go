package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/aacebo/equinox-api/src/api"
	"github.com/aacebo/equinox-api/src/db"
	"github.com/aacebo/equinox-api/src/log"
)

var _log = log.New("main")

func main() {
	var env = os.Getenv("GO_ENV")
	var err = godotenv.Load(fmt.Sprintf(".env.%s", env))

	if err != nil {
		_log.Error(err)
	}

	var port = os.Getenv("PORT")

	_log.Infof("running on %s", env)

	var db, dberr = db.Connect()

	if dberr != nil {
		_log.Error(dberr)
	}

	_log.Info("connected to database...")

	defer db.Close()

	var router, routererr = api.Router(env, db)

	if routererr != nil {
		_log.Error(routererr)
	}

	router.Run(fmt.Sprintf("localhost:%s", port))
}
