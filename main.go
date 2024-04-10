package main

import (
	"fmt"
	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Loading configuration file ...")
	config.ConnectToDB()
	r := gin.Default()
	// r.Use(config.ApiMiddleware())
	r.Use(cors.Default())
	r.MaxMultipartMemory = 8 << 20
	r.Static("/assets", "./assets")
	routers.SetupUserRoutes(r)
	routers.SetupImageRoutes(r, cfg)
	routers.SetupTextRoutes(r, cfg)
	routers.SetupPromptRoutes(r)
	r.Run(":8080")
}
