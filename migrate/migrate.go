package main

import (
	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/models"
)

func init() {
	config.LoadConfig()
	config.ConnectToDB()
}

func main() {
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Prompt{})
	config.DB.AutoMigrate(&models.History{})
}
