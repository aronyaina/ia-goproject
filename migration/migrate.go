package main

import (
	"fmt"
	"os"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/models"
)

func init() {
	config.LoadConfig()
	config.ConnectToDB()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing argument: Please provide a command (up or down).")
		return
	}

	command := os.Args[1]

	switch command {
	case "up":
		fmt.Println("Migrating User")
		config.DB.AutoMigrate(&models.User{})
		fmt.Println("Migrating Prompt")
		config.DB.AutoMigrate(&models.Prompt{})
	case "down":
		fmt.Println("Dropping User table (if exists)")
		config.DB.Migrator().DropTable(&models.User{})
		fmt.Println("Dropping Prompt table (if exists)")
		config.DB.Migrator().DropTable(&models.Prompt{})
	case "restart":
		fmt.Println("Migrating User")
		config.DB.AutoMigrate(&models.User{})
		fmt.Println("Migrating Prompt")
		config.DB.AutoMigrate(&models.Prompt{})
		fmt.Println("Dropping User table (if exists)")
		config.DB.Migrator().DropTable(&models.User{})
		fmt.Println("Dropping Prompt table (if exists)")
		config.DB.Migrator().DropTable(&models.Prompt{})
	default:
		fmt.Println("Invalid command. Please use 'up' or 'down'.")
	}
}
