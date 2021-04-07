package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/src/page"
)

func NotFound(ctx *gin.Context) {
	ctx.JSON(
		http.StatusNotFound,
		gin.H{"error": NewError(http.StatusNotFound)},
	)
}

func BadRequest(ctx *gin.Context) {
	ctx.JSON(
		http.StatusBadRequest,
		gin.H{"error": NewError(http.StatusBadRequest)},
	)
}

func Ok(ctx *gin.Context, data interface{}) {
	ctx.JSON(
		http.StatusOK,
		New(data),
	)
}

func OkPaged(ctx *gin.Context, page *page.Page, total int, data interface{}) {
	ctx.JSON(
		http.StatusOK,
		NewPaged(
			page,
			total,
			data,
		),
	)
}
