package routers

import (
	"github.com/aronyaina/ia-goproject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userRouter := router.Group("/users")
	{
		userRouter.GET("/", controllers.GetAllUser)
		userRouter.POST("/", controllers.CreateUser)
		userRouter.GET("/:id", controllers.GetUserById)
	}
}
