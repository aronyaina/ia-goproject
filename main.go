package main

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/services"
	"github.com/disintegration/imaging"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configuration, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.Use(cors.Default())
	r.MaxMultipartMemory = 8 << 20

	r.POST("/text-to-images/generate", func(c *gin.Context) {
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
		err = imaging.Save(image, outputFilename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Image created successfully", "filename": outputFilename})
	})

	r.POST("/image-to-text", func(c *gin.Context) {
		dirName := services.ImageUploader(c)
		output, err := services.ImageToText(dirName, configuration.Server.ImageToText, configuration.Server.Token)
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
		output, err := services.ImageToText(filename, configuration.Server.ImageToText, configuration.Server.Token)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": output,
		})
	})

	r.POST("/text-to-texts/generate", func(c *gin.Context) {
		var payload services.Payload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err := services.TextToText(payload, configuration.Server.TextSummerization, configuration.Server.Token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	r.POST("/text-classifications", func(c *gin.Context) {
		var payload services.Payload
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response, err := services.TextToText(payload, configuration.Server.TextClassification, configuration.Server.Token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response)
	})

	r.Run(":8080")
}
