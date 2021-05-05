package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/aacebo/equinox-api/api"
	"github.com/aacebo/equinox-api/db"
	"github.com/aacebo/equinox-api/log"
)

func main() {
	var env = os.Getenv("GO_ENV")
	var err = godotenv.Load(fmt.Sprintf(".env.%s", env))

	if err != nil {
		log.Error.Fatal(err)
	}

	var port = os.Getenv("PORT")

	log.Info.Printf("running on %s", env)

	var db = db.Connect()
	defer db.Close()

	log.Info.Print("connected to database...")

	api.Router(env, db).Run(fmt.Sprintf("localhost:%s", port))
}
