package controllers

import (
	"net/http"

	"github.com/aronyaina/ia-goproject/services"
	"github.com/gin-gonic/gin"
)

func GenerateText(c *gin.Context) {
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
	//CreatePrompt(c, "TEXT_GENERATION", response, c.Param("user_id"))
	c.JSON(http.StatusOK, response)
}

func TextClassification(c *gin.Context) {
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
	//CreatePrompt(c, "TEXT_CLASSIFICAITON", response, c.Param("user_id"))
	c.JSON(http.StatusOK, response)

}
