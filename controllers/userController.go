package controllers

import (
	"net/http"
	"time"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserPayload struct {
	Email string `json:"email"`
}

func CreateUser(c *gin.Context) {
	var payload UserPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		ID:        uuid.New().String(),
		Email:     payload.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	config.ConnectToDB()
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, user)
	}
}

func GetAllUser(c *gin.Context) {
	config.ConnectToDB()
	var users []models.User
	result := config.DB.Find(&users)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, users)
	}
}

func GetUserById(c *gin.Context) {
	config.ConnectToDB()
	var user models.User

	result := config.DB.Where("ID =?", c.Param("id")).First(&user)

	if result.Error != nil {
		if result.RowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": result.Error.Error()})
	}
	c.JSON(200, user)
}
