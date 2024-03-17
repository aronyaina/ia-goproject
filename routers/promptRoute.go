package routers

import (
	"github.com/aronyaina/ia-goproject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupPromptRoutes(router *gin.Engine) {
	promptRouter := router.Group("/prompts")
	{
		promptRouter.GET("/:user_id", controllers.GetAllPromptByUserId)
		promptRouter.DELETE("/", controllers.DeleteOnePromptById)
	}
}
