package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/aacebo/equinox-api/src/api/organizations"
	"github.com/aacebo/equinox-api/src/db"
	"github.com/aacebo/equinox-api/src/logger"
)

func main() {
	var env = os.Getenv("GO_ENV")
	var err = godotenv.Load(fmt.Sprintf(".env.%s", env))
	var log = logger.New("seed")

	if err != nil {
		log.Error(err)
	}

	log.Infof("running on %s", env)

	var db = db.Connect()
	defer db.Close()

	log.Info("connected to database...")

	var orgr = organizations.NewRepository(db)
	var orgs = organizations.NewSeed()

	log.Infof("%d organizations", len(orgs.Organizations))

	for _, o := range orgs.Organizations {
		orgr.Upsert(
			o.ID,
			o.Slug,
			o.Name,
			o.CreatedAt,
			o.UpdatedAt,
		)
	}
}
