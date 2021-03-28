package organization

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Find(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "hello world"})
}
