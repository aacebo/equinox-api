package organizations

import (
	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/src/logger"
)

var log = logger.New("organizations")

func Controller(orgr *Repository) func(r *gin.Engine) *gin.RouterGroup {
	return func(r *gin.Engine) *gin.RouterGroup {
		var g = r.Group("/organizations")
		{
			g.GET("/", Find(orgr))
			g.GET("/:org_slug", FindOne(orgr))
		}

		return g
	}
}
