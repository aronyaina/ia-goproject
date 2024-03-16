package main

import (
	"log"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var configuration *config.Config

func init() {
	var err error
	configuration, err = config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	config.ConnectToDB()
	r := gin.Default()
	// r.Use(config.ApiMiddleware())
	r.Use(cors.Default())
	r.MaxMultipartMemory = 8 << 20

	routers.SetupUserRoutes(r)
	routers.SetupImageRoutes(r, configuration)
	routers.SetupTextRoutes(r)
	r.Run(":8080")
}
