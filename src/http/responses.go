package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		gin.H{"data": data},
	)
}
