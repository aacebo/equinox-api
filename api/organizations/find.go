package organizations

import (
	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/page"
	"github.com/aacebo/equinox-api/response"
)

func Find(orgr *Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var page = page.New(ctx)

		if page == nil {
			response.BadRequest(ctx)
			return
		}

		var orgs, total = orgr.Find(page)

		response.OkPaged(ctx, page, total, orgs)
	}
}
