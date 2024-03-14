package main

import (
	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(config.ApiMiddleware())
	r.Use(cors.Default())
	r.MaxMultipartMemory = 8 << 20

	routers.SetupUserRoutes(r)
	routers.SetupImageRoutes(r)
	routers.SetupTextRoutes(r)
	r.Run(":8080")
}
