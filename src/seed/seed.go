package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/aacebo/equinox-api/src/api/organizations"
	"github.com/aacebo/equinox-api/src/db"
	"github.com/aacebo/equinox-api/src/logger"
)

func main() {
	var env = os.Getenv("GO_ENV")
	var count, _ = strconv.Atoi(os.Getenv("SEED_COUNT"))
	var err = godotenv.Load(fmt.Sprintf(".env.%s", env))
	var log = logger.New("seed")

	if err != nil {
		log.Error(err)
	}

	log.Infof("running on %s", env)

	var db, dberr = db.Connect()

	if dberr != nil {
		log.Error(dberr)
	}

	log.Info("connected to database...")

	defer db.Close()

	var orgr = organizations.NewRepository(db)

	for i := 0; i < count; i++ {
		orgr.Mock()
	}
}
