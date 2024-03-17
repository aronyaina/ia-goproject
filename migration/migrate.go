package main

import (
	"fmt"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/models"
)

func init() {
	config.LoadConfig()
	config.ConnectToDB()
}

func main() {
	fmt.Println("Migrating User")
	config.DB.AutoMigrate(&models.User{})
	fmt.Println("Migrating Prompt")
	config.DB.AutoMigrate(&models.Prompt{})
}
