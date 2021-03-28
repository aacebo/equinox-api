package api

import (
	"github.com/aacebo/equinox-api/src/api/organization"
	"github.com/gin-gonic/gin"
)

func Router(env string) *gin.Engine {
	var mode = gin.DebugMode

	if env == "production" {
		mode = gin.ReleaseMode
	}

	gin.SetMode(mode)
	var r = gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	organization.Controller(r)

	return r
}
