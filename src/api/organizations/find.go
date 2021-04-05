package organizations

import (
	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/src/page"
	"github.com/aacebo/equinox-api/src/response"
)

func Find(orgr *Repository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var orgs, total = orgr.Find(page.New(ctx))

		response.OkPaged(ctx, total, orgs)
	}
}
