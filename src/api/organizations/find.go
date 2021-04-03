package organizations

import (
	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/src/http"
	"github.com/aacebo/equinox-api/src/page"
)

func Find(orgr *Repository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var orgs = orgr.Find(page.New(ctx))

		http.Ok(ctx, orgs)
	}
}
