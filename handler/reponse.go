package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message":    message,
		"statusCode": code,
	})
}

func SendSuccess(c *gin.Context, operation string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfully", operation),
		"data":    data,
	})
}
