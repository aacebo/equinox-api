package organizations

import (
	"github.com/gin-gonic/gin"
)

func Controller(orgr Repository) func(r *gin.Engine) *gin.RouterGroup {
	return func(r *gin.Engine) *gin.RouterGroup {
		var g = r.Group("/organizations")
		{
			g.GET("/", Find(orgr))
		}

		return g
	}
}
