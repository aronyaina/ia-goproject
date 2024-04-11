package routers

import (
	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTextRoutes(router *gin.Engine, config *config.Config) {
	textRouter := router.Group("/texts")
	{
		textRouter.POST("/generation", func(c *gin.Context) {
			controllers.TextGeneration(c, config)
		})

		textRouter.POST("/summerization", func(c *gin.Context) {
			controllers.TextSummerization(c, config)
		})
		textRouter.POST("/classification", func(c *gin.Context) {
			controllers.TextClassification(c, config)
		})

	}
}
