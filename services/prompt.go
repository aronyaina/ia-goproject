package services

import (
	"fmt"
	"time"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreatePrompt(c *gin.Context, tag string, input string, result string, user_id string) {
	prompt := models.Prompt{
		ID:        uuid.New().String(),
		Tag:       tag,
		Result:    result,
		Input:     input,
		UserID:    user_id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	config.ConnectToDB()
	response := config.DB.Create(&prompt)
	fmt.Println("Prompt data created successfully")

	if response.Error != nil {
		c.JSON(500, gin.H{"error": response.Error.Error()})
	}
}
