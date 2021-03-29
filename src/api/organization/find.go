package organization

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Find(orgRepository Repository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var orgs = orgRepository.Find()
		ctx.JSON(http.StatusOK, gin.H{"data": orgs})
	}
}
