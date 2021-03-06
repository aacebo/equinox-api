package organizations

import (
	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/response"
)

func FindOne(orgr *Repository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var org = orgr.FindBySlug(ctx.Param("org_slug"))

		if org == nil {
			response.NotFound(ctx)
			return
		}

		response.Ok(ctx, org)
	}
}
