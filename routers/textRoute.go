package routers

import (
	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTextRoutes(router *gin.Engine, config *config.Config) {
	textRouter := router.Group("/texts")
	{
		textRouter.POST("/text/:user_id", func(c *gin.Context) {
			controllers.GenerateText(c, config)
		})
		textRouter.POST("/classification/:user_id", func(c *gin.Context) {
			controllers.TextClassification(c, config)
		})

	}
}
