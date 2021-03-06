package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"github.com/aacebo/equinox-api/api/organizations"
	"github.com/aacebo/equinox-api/db"
	"github.com/aacebo/equinox-api/log"
)

func main() {
	var env = os.Getenv("GO_ENV")
	var err = godotenv.Load(fmt.Sprintf(".env.%s", env))

	if err != nil {
		log.Error.Fatal(err)
	}

	log.Info.Printf("running on %s", env)

	var db = db.Connect()
	defer db.Close()

	log.Info.Print("connected to database...")

	var orgr = organizations.NewRepository(db)
	var orgs = organizations.NewSeed()

	log.Info.Printf("%d organizations", len(orgs.Organizations))

	for _, o := range orgs.Organizations {
		orgr.Upsert(
			o.ID,
			o.Key,
			o.Slug,
			o.Name,
			*o.CreatedAt,
			*o.UpdatedAt,
		)
	}
}
