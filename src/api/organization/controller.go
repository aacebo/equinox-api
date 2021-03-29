package organization

import (
	"github.com/gin-gonic/gin"
)

func Controller(orgRepository Repository) func(r *gin.Engine) *gin.RouterGroup {
	return func(r *gin.Engine) *gin.RouterGroup {
		var g = r.Group("/organizations")
		{
			g.GET("/", Find(orgRepository))
		}

		return g
	}
}
