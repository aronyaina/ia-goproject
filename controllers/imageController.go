package controllers

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/services"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
)

var configuration *config.Config

func init() {
	_, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	config.ConnectToDB()
}
func ImageToText(c *gin.Context) {
	dirName := services.ImageUploader(c)
	output, err := services.ImageToText(dirName, configuration.Server.ImageToText, configuration.Server.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//CreatePrompt(c, "IMAGE_TO_TEXT", output, c.Param("user_id"))

	c.JSON(http.StatusOK, gin.H{
		"message": output,
	})
}

func TextToImage(c *gin.Context) {
	var payload services.Payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imageBytes, err := services.TextToImageQuery(configuration.Server.TextToImage, configuration.Server.Token, payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	image, err := imaging.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Type", "image/jpeg")
	outputFilename := "output.jpg"
	//CreatePrompt(c, "TEXT_TO_IMAGE", outputFilename, c.Param("user_id"))
	err = imaging.Save(image, outputFilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image created successfully", "filename": outputFilename})
}

func ImageClassification(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	filename := filepath.Base(file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	output, err := services.ImageToText(filename, configuration.Server.ImageToText, configuration.Server.Token)
	if err != nil {
		log.Fatal(err)
	}

	//CreatePrompt(c, "IMAGE_CLASSIFICATION", output, c.Param("user_id"))
	c.JSON(http.StatusOK, gin.H{
		"message": output,
	})
}
