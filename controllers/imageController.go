package controllers

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/services"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ImageToText(c *gin.Context, config *config.Config) {
	if config == nil {
		handleError(c, errors.New("No configuration provided."), http.StatusBadRequest)
		return
	}
	dirName := services.ImageUploader(c)
	output, err := services.ImageToText(dirName, config.Server.ImageToText, config.Server.Token)
	if err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}

	var text string
	if len(output) > 0 && output[0]["generated_text"] != nil {
		text = output[0]["generated_text"].(string)
	}

	done := make(chan bool)
	go func() {
		services.CreatePrompt(c, "IMAGE_TO_TEXT", text, dirName, c.Param("user_id"))
		done <- true
	}()
	<-done

	c.JSON(http.StatusOK, output[0])
}

func TextToImage(c *gin.Context, config *config.Config) {
	if config == nil {
		handleError(c, errors.New("No configuration provided."), http.StatusBadRequest)
		return
	}
	var payload services.Payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}

	imageBytes, err := services.TextToImageQuery(config.Server.TextToImage, config.Server.Token, payload)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}

	image, err := imaging.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		handleError(c, err, http.StatusInternalServerError)
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
	services.CreatePrompt(c, "TEXT_TO_IMAGE", "assets/"+outputFilename, "assets/"+outputFilename, c.Param("user_id"))
	c.JSON(http.StatusOK, imageData)
}

func ImageClassification(c *gin.Context, config *config.Config) {
	if config == nil {
		handleError(c, errors.New("No configuration provided."), http.StatusBadRequest)
		return
	}
	dirName := services.ImageUploader(c)

	output, err := services.ImageToText(dirName, config.Server.ImageClassification, config.Server.Token)
	if err != nil {
		handleError(c, err, http.StatusBadRequest)
	}

	bestLabel := output[0]["label"].(string)
	done := make(chan bool)
	go func() {
		services.CreatePrompt(c, "IMAGE_CLASSIFICATION", bestLabel, dirName, c.Param("user_id"))
		done <- true
	}()
	<-done

	c.JSON(http.StatusOK, output)
}

func handleError(c *gin.Context, err error, statusCode int) {
	fmt.Println(err)
	c.JSON(statusCode, gin.H{"error": err.Error()})
}
