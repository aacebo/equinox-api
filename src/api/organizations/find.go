package organizations

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/src/errors"
)

func Find(orgRepository Repository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var orgs, err = orgRepository.Find()

		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": errors.New(http.StatusInternalServerError, err)},
			)

			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": orgs})
	}
}
