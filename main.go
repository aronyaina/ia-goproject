package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Payload struct {
	Inputs string `json:"inputs"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running on port 8080",
		})
	})

	r.POST("/text-to-images/generate", func(c *gin.Context) {
		var payload Payload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		imageBytes, err := textToImageQuery(os.Getenv("URL_TEXT_TO_IMAGE"), os.Getenv("API_TOKEN"), payload)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		image, err := imaging.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		outputFilename := "output.jpg"
		err = imaging.Save(image, outputFilename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Image created successfully", "filename": outputFilename})
	})
	r.GET("/image-to-text/results", func(c *gin.Context) {
		output, err := ImageToTextQuery("output.jpg", os.Getenv("URL_IMAGE_TO_TEXT"), os.Getenv("API_TOKEN"))
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": output,
		})
	})
	r.GET("/image-classification", func(c *gin.Context) {
		output, err := ImageToTextQuery("output.jpg", os.Getenv("URL_IMAGE_CLASSIFICATION"), os.Getenv("API_TOKEN"))
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": output,
		})

	})
	r.POST("/text-to-texts/generate", func(c *gin.Context) {
		var payload Payload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err := textToText(payload, os.Getenv("URL_TEXT_SUMMERIZATION"), os.Getenv("API_TOKEN"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	})
	r.POST("/text-classifications", func(c *gin.Context) {
		var payload Payload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err := textToText(payload, os.Getenv("URL_TEXT_CLASSIFICATION"), os.Getenv("API_TOKEN"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	})
	r.Run(":8080")
}
