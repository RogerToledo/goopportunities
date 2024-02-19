package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}
