package routers

import (
	"github.com/aronyaina/ia-goproject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupHistoryRoutes(router *gin.Engine) {
	historyRouter := router.Group("/histories")
	{
		historyRouter.GET("/:user_id", controllers.GetAllHistoryByUserID)
		historyRouter.POST("/", controllers.CreateHistory)
		historyRouter.DELETE("/:id", controllers.DeleteHistoryById)
	}
}
