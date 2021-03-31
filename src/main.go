package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

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

	log.Printf("running on %s", env)

	var db = db.Connect()

	log.Println("connected to database...")

	defer db.Close()

	var router = api.Router(env, db)

	router.Run(fmt.Sprintf("localhost:%s", port))
}
