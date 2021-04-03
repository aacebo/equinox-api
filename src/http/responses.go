package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(
		http.StatusInternalServerError,
		gin.H{"error": NewError(http.StatusInternalServerError, err)},
	)
}

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
