package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/aacebo/equinox-api/src/api"
)

func main() {
	env := os.Getenv("GO_ENV")
	envErr := godotenv.Load(fmt.Sprintf(".env.%s", env))

	if envErr != nil {
		log.Fatal(envErr)
	}

	port := os.Getenv("PORT")

	log.Infof("running on %s", env)
	log.Infof("listening on port %s...", port)

	httpErr := http.ListenAndServe(fmt.Sprintf("localhost:%s", port), api.Router())

	if httpErr != nil {
		log.Fatal(httpErr)
	}
}
