package router

import (
	"github.com/gin-gonic/gin"
	"github.com/me/goopportunities/docs"
	"github.com/me/goopportunities/handler"
	"github.com/me/goopportunities/handler/opening"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	const basePath = "/api/v1"
	handler.InitializeHandler()

	v1 := router.Group(basePath)

	v1.GET("opening", opening.ShowOpening)
	v1.POST("opening", opening.CreateOpening)
	v1.DELETE("opening", opening.DeleteOpening)
	v1.PUT("opening", opening.UpdateOpening)
	v1.GET("openings", opening.ListOpening)

	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
