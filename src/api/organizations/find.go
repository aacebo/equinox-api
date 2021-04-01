package organizations

import (
	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/src/http"
)

func Find(orgr Repository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var orgs, err = orgr.Find()

		if err != nil {
			http.InternalServerError(ctx, err)
			return
		}

		http.Ok(ctx, orgs)
	}
}
