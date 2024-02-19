package router

import (
	"github.com/gin-gonic/gin"
	"github.com/me/goopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	
	v1.GET("opening", handler.ShowOpening)

	v1.POST("opening", handler.CreateOpening)

	v1.DELETE("opening", handler.DeleteOpening)

	v1.PUT("opening", handler.UpdateOpening)

	v1.GET("opening", handler.ListOpening)	
}	