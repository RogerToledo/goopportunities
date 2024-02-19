package opening

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "PUT Opening",
	})
}
