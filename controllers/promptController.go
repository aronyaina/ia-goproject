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

func CreatePrompt(c *gin.Context, payload PromptPayload) {
	prompt := models.Prompt{
		ID:        uuid.New().String(),
		Tag:       payload.Tag,
		Result:    payload.Result,
		UserID:    payload.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	config.ConnectToDB()
	result := config.DB.Create(&prompt)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, gin.H{"User": prompt})
	}
}

func GetAllPromptByUserId(c *gin.Context, user_id string) {
	config.ConnectToDB()
	var prompts models.Prompt
	result := config.DB.Where("UserID =?", user_id).Find(&prompts)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, gin.H{"Prompts": prompts})
	}
}

func GetOnePromptById(c *gin.Context, id string) {
	config.ConnectToDB()
	var prompt models.Prompt
	result := config.DB.Where("ID =?", id).First(&prompt)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, gin.H{"User": prompt})
	}
}

func deleteOnePromptById(c *gin.Context, id string) {
	config.ConnectToDB()
	var prompt models.Prompt
	result := config.DB.Where("ID =?", id).Delete(&prompt)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, gin.H{"User": prompt})
	}
}
