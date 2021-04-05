package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aacebo/equinox-api/src/page"
	"github.com/aacebo/equinox-api/src/response"
)

func NotFound(ctx *gin.Context) {
	ctx.JSON(
		http.StatusNotFound,
		gin.H{"error": NewError(http.StatusNotFound)},
	)
}

func Ok(ctx *gin.Context, data interface{}) {
	ctx.JSON(
		http.StatusOK,
		response.New(data),
	)
}

func OkPaged(ctx *gin.Context, total int, data interface{}) {
	ctx.JSON(
		http.StatusOK,
		response.NewPaged(
			page.New(ctx),
			total,
			data,
		),
	)
}
