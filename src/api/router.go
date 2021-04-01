package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/src/api/organizations"
)

func Router(env string, db *sql.DB) (*gin.Engine, error) {
	var mode = gin.DebugMode

	if env == "production" {
		mode = gin.ReleaseMode
	}

	gin.SetMode(mode)

	var r = gin.New()
	var orgr, orgrerr = organizations.NewRepository(db)

	if orgrerr != nil {
		return nil, orgrerr
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	organizations.Controller(*orgr)(r)

	return r, nil
}
