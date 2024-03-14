package routers

import (
	"github.com/aronyaina/ia-goproject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupTextRoutes(router *gin.Engine) {
	textRouter := router.Group("/texts")
	{
		textRouter.POST("/text/:user_id", controllers.GenerateText)
		textRouter.POST("/classification/:user_id", controllers.TextClassification)
	}
}
