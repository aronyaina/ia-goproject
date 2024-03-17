package controllers

import (
	"net/http"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/models"
	"github.com/gin-gonic/gin"
)

type PromptPayload struct {
	Tag    string `json:"tag"`
	Result string `json:"result"`
	UserID string `json:"user_id"`
}

func GetAllPromptByUserId(c *gin.Context) {
	config.ConnectToDB()
	var prompts []models.Prompt
	result := config.DB.Where("user_id =?", c.Param("user_id")).Find(&prompts)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": result.Error.Error()})
	}
	c.JSON(200, prompts)
}

func DeleteOnePromptById(c *gin.Context) {
	config.ConnectToDB()
	var prompt models.Prompt
	result := config.DB.Where("ID =?", c.Param("id")).Delete(&prompt)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, prompt)
	}
}
