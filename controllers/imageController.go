package controllers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/services"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var configuration *config.Config

func ImageToText(c *gin.Context, config *config.Config) {
	if config != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("No configuration provided.")})
		return
	}
	dirName := services.ImageUploader(c)
	output, err := services.ImageToText(dirName, configuration.Server.ImageToText, configuration.Server.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var text string
	if len(output) > 0 && output[0]["generated_text"] != nil {
		text = output[0]["generated_text"].(string)
	}

	done := make(chan bool)
	go func() {
		CreatePrompt(c, "IMAGE_TO_TEXT", text, dirName, c.Param("user_id"))
		done <- true
	}()
	<-done

	fmt.Println("0. Test Passed")

	c.JSON(http.StatusOK, output)
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
	id := uuid.New()
	outputFilename := id.String()
	err = imaging.Save(image, outputFilename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	imageData := base64.StdEncoding.EncodeToString(imageBytes)
	CreatePrompt(c, "TEXT_TO_IMAGE", "assets/"+outputFilename, "assets/"+outputFilename, c.Param("user_id"))
	c.JSON(http.StatusOK, imageData)
}

func ImageClassification(c *gin.Context) {
	configuration, err := config.LoadConfig()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dirName := services.ImageUploader(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	output, err := services.ImageToText(dirName, configuration.Server.ImageClassification, configuration.Server.Token)
	if err != nil {
		log.Fatal(err)
	}

	bestLabel := output[0]["label"].(string)
	done := make(chan bool)
	go func() {
		CreatePrompt(c, "IMAGE_CLASSIFICATION", bestLabel, dirName, c.Param("user_id"))
		done <- true
	}()
	<-done

	c.JSON(http.StatusOK, output)
}
