package controllers

import (
	"time"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context, payload UserPayload) {
	user := models.User{
		ID:        uuid.New().String(),
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  payload.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	config.ConnectToDB()
	result := config.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, gin.H{"User": user})
	}
}

func GetAllUser(c *gin.Context) {
	config.ConnectToDB()
	var users []models.User
	result := config.DB.Find(&users)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, gin.H{"Users": users})
	}
}

func GetUserById(c *gin.Context, id string) {
	config.ConnectToDB()
	var user models.User
	result := config.DB.Where("ID =?", id).First(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
	} else {
		c.JSON(200, gin.H{"User": user})
	}
}
