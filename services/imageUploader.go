package services

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ImageUploader(c *gin.Context) (dir_name string) {
	image, err := c.FormFile("image")
	fileType := strings.Split(image.Filename, ".")[1]
	fmt.Println(fileType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	log.Println(image.Filename)
	id := uuid.New()
	fileName := "assets/" + id.String() + "." + fileType
	c.SaveUploadedFile(image, fileName)
	return fileName
}
