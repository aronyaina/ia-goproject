package routers

import (
	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupImageRoutes(router *gin.Engine, config *config.Config) {
	imageRouter := router.Group("/images")
	{
		imageRouter.POST("/text/:user_id", func(c *gin.Context) {
			controllers.ImageToText(c, config)
		})
		imageRouter.POST("/image/:user_id", func(c *gin.Context) {
			controllers.TextToImage(c, config)
		})
		imageRouter.POST("/classification/:user_id", func(c *gin.Context) {
			controllers.ImageClassification(c, config)
		})
	}
}
