package controllers

import (
	"time"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PromptPayload struct {
	Tag    string `json:"tag"`
	Result string `json:"result"`
	UserID string `json:"user_id"`
}

func CreatePrompt(c *gin.Context, tag string, result string, image_name string, user_id string) {
	prompt := models.Prompt{
		ID:        uuid.New().String(),
		Tag:       tag,
		Result:    result,
		ImageName: image_name,
		UserID:    user_id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	config.ConnectToDB()
	response := config.DB.Create(&prompt)
	if response.Error != nil {
		c.JSON(500, gin.H{"error": response.Error.Error()})
	}
}

func GetAllPromptByUserId(c *gin.Context) {
	config.ConnectToDB()
	var prompts models.Prompt
	result := config.DB.Where("UserID =?", c.Param("user_id")).Find(&prompts)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, prompts)
	}
}

func deleteOnePromptById(c *gin.Context) {
	config.ConnectToDB()
	var prompt models.Prompt
	result := config.DB.Where("ID =?", c.Param("id")).Delete(&prompt)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, prompt)
	}
}
