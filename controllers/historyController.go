package controllers

import (
	"time"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type historyPayload struct {
	Title    string `json:"title"`
	UserId   string `json:"user_id"`
	PromptId string `json:"prompt_id"`
}

func CreateHistory(c *gin.Context, payload historyPayload) {
	history := models.History{
		ID:        uuid.New().String(),
		Title:     payload.Title,
		UserID:    payload.UserId,
		PromptID:  payload.PromptId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	config.ConnectToDB()
	result := config.DB.Create(&history)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, gin.H{"User": history})
	}
}

func DeleteHistoryById(c *gin.Context, id string) {
	config.ConnectToDB()
	var history models.History
	result := config.DB.Where("ID =?", id).Delete(&history)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, gin.H{"User": history})
	}
}

func GetAllHistoryByUserID(c *gin.Context, user_id string) {
	config.ConnectToDB()
	var history models.Prompt
	result := config.DB.Where("UserID =?", user_id).Find(&history)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, gin.H{"Prompts": history})
	}
}
