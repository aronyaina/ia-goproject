package routers

import (
	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupImageRoutes(router *gin.Engine, config *config.Config) {
	imageRouter := router.Group("/images")
	{
		imageRouter.POST("/text", func(c *gin.Context) {
			controllers.ImageToText(c, config)
		})
		imageRouter.POST("/generation", func(c *gin.Context) {
			controllers.TextToImage(c, config)
		})
		imageRouter.POST("/classification", func(c *gin.Context) {
			controllers.ImageClassification(c, config)
		})
	}
}
