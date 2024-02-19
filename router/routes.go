package router

import (
	"github.com/gin-gonic/gin"
	"github.com/me/goopportunities/handler"
	"github.com/me/goopportunities/handler/opening"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	
	v1.GET("opening", opening.ShowOpening)

	v1.POST("opening", opening.CreateOpening)

	v1.DELETE("opening", opening.DeleteOpening)

	v1.PUT("opening", opening.UpdateOpening)

	v1.GET("opening", opening.ListOpening)	
}	