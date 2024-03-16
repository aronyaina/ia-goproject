package controllers

import (
	"net/http"
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

func CreateHistory(c *gin.Context) {
	var payload historyPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
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
		c.JSON(200, history)
	}
}

func DeleteHistoryById(c *gin.Context) {
	config.ConnectToDB()
	var history models.History
	result := config.DB.Where("ID =?", c.Param("id")).Delete(&history)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, history)
	}
}

func GetAllHistoryByUserID(c *gin.Context) {
	config.ConnectToDB()
	var history models.Prompt
	result := config.DB.Where("UserID =?", c.Param("user_id")).Find(&history)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, history)
	}
}
