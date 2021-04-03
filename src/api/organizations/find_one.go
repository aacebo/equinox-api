package organizations

import (
	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/src/http"
)

func FindOne(orgr Repository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var org, err = orgr.FindBySlug(ctx.Param("org_slug"))

		if err != nil {
			http.InternalServerError(ctx, err)
			return
		}

		if org == nil {
			http.NotFound(ctx, err)
			return
		}

		http.Ok(ctx, org)
	}
}
