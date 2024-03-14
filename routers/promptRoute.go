package routers

import (
	"github.com/aronyaina/ia-goproject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupPromptRoutes(router *gin.Engine) {
	promptRouter := router.Group("/prompts")
	{
		promptRouter.GET("/", controllers.GetAllPromptByUserId)
		promptRouter.GET("/:id", controllers.GetOnePromptById)
	}
}
