package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/services"
	"github.com/disintegration/imaging"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Payload struct {
	Inputs string `json:"inputs"`
}

func init() {
	configuration, err = config.LoadConfig()
	if err != nil {
		panic(err)
	}
}
func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.MaxMultipartMemory = 8 << 20

	r.POST("/text-to-images/generate", func(c *gin.Context) {
		var payload Payload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		imageBytes, err := services.TextToImageQuery(os.Getenv("URL_TEXT_TO_IMAGE"), os.Getenv("API_TOKEN"), payload)
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
		err = imaging.Save(image, outputFilename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Image created successfully", "filename": outputFilename})
	})

	r.POST("/image-to-text", func(c *gin.Context) {
		dirName := services.ImageUploader(c)
		output, err := services.ImageToText(dirName, os.Getenv("URL_IMAGE_TO_TEXT"), os.Getenv("API_TOKEN"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": output,
		})
	})

	r.GET("/image-classification", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		filename := filepath.Base(file.Filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		output, err := services.ImageToText(filename, os.Getenv("URL_IMAGE_CLASSIFICATION"), os.Getenv("API_TOKEN"))
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
		response, err := services.TextToText(payload, os.Getenv("URL_TEXT_SUMMERIZATION"), os.Getenv("API_TOKEN"))
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
		response, err := services.TextToText(payload, os.Getenv("URL_TEXT_CLASSIFICATION"), os.Getenv("API_TOKEN"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	r.Run(":8080")
}
