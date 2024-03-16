package routers

import (
	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupImageRoutes(router *gin.Engine, config *config.Config) {
	imageRouter := router.Group("/images")
	{
		imageRouter.POST("/text/:user_id", controllers.ImageToText)
		imageRouter.POST("/image/:user_id", controllers.TextToImage)
		imageRouter.POST("/classification/:user_id", controllers.ImageClassification)
	}
}
