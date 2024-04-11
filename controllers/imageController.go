package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ImageToText(c *gin.Context, config *config.Config) {
	if config == nil {
		handleError(c, errors.New("no configuration provided"), http.StatusBadRequest)
		return
	}
	dirName := services.ImageUploader(c)
	output, err := services.ImageToText(dirName, config.Server.ImageToText, config.Server.Token)
	if err != nil {
		handleError(c, err, http.StatusBadRequest)
		return
	}

	//var text string
	//if len(output) > 0 && output[0]["generated_text"] != nil {
	//	text = output[0]["generated_text"].(string)
	//}

	//done := make(chan bool)
	//go func() {
	//	services.CreatePrompt(c, "IMAGE_TO_TEXT", dirName, text, c.Param("user_id"))
	//	done <- true
	//}()
	//<-done

	c.JSON(http.StatusOK, output[0])
}

func TextToImage(c *gin.Context, config *config.Config) {
	if config == nil {
		handleError(c, errors.New("no configuration provided"), http.StatusBadRequest)
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

	outputFilename, err := SaveImage(imageBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"image_url": outputFilename})
	//services.CreatePrompt(c, "TEXT_TO_IMAGE", payload.Inputs, "assets/"+outputFilename, c.Param("user_id"))
	//c.JSON(http.StatusOK, imageData)
}

func ImageClassification(c *gin.Context, config *config.Config) {
	if config == nil {
		handleError(c, errors.New("no configuration provided"), http.StatusBadRequest)
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
		services.CreatePrompt(c, "IMAGE_CLASSIFICATION", dirName, bestLabel, c.Param("user_id"))
		done <- true
	}()
	<-done

	c.JSON(http.StatusOK, output)
}

func handleError(c *gin.Context, err error, statusCode int) {
	fmt.Println(err)
	c.JSON(statusCode, gin.H{"error": err.Error()})
}

func SaveImage(imageBytes []byte) (outputFilename string, err error) {
	id := uuid.New()
	filename := "assets/" + id.String() + ".jpeg"
	error := os.WriteFile(filename, imageBytes, 0644)
	if error != nil {
		return "", error
	}
	return filename, nil
}
