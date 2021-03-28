package organization

import (
	"github.com/gin-gonic/gin"
)

func Controller(r *gin.Engine) *gin.RouterGroup {
	var g = r.Group("/organizations")
	{
		g.GET("/", Find)
	}

	return g
}
