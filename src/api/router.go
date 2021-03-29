package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/src/api/organization"
)

func Router(env string, db *sql.DB) *gin.Engine {
	var mode = gin.DebugMode

	if env == "production" {
		mode = gin.ReleaseMode
	}

	gin.SetMode(mode)
	var r = gin.New()
	var orgRepository = organization.NewRepository(db)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	organization.Controller(*orgRepository)(r)

	return r
}
