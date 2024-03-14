package routers

import (
	"github.com/aronyaina/ia-goproject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupImageRoutes(router *gin.Engine) {
	imageRouter := router.Group("/images")
	{
		imageRouter.POST("/text/:user_id", controllers.TextToImage)
		imageRouter.POST("/image/:user_id", controllers.ImageToText)
		imageRouter.POST("/classification/:user_id", controllers.ImageClassification)
	}
}
