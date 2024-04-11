package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/aronyaina/ia-goproject/config"
	"github.com/aronyaina/ia-goproject/services"
	"github.com/gin-gonic/gin"
)

func TextGeneration(c *gin.Context, config *config.Config) {
	var payload services.Payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var response []interface{}
	response, err := services.TextToText(payload, config.Server.TextGeneration, config.Server.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//slice, ok := response[0].(map[string]interface{})
	//if !ok {
	//	fmt.Println("Error while processing response")
	//	return
	//}

	//services.CreatePrompt(c, "TEXT_GENERATION", payload.Inputs, slice["summary_text"].(string), c.Param("user_id"))
	c.JSON(http.StatusOK, response[0])
}
func TextSummerization(c *gin.Context, config *config.Config) {
	var payload services.Payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var response []interface{}
	response, err := services.TextToText(payload, config.Server.TextSummerization, config.Server.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//slice, ok := response[0].(map[string]interface{})
	//if !ok {
	//	fmt.Println("Error while processing response")
	//	return
	//}

	//services.CreatePrompt(c, "TEXT_GENERATION", payload.Inputs, slice["summary_text"].(string), c.Param("user_id"))
	c.JSON(http.StatusOK, response[0])
}

func TextClassification(c *gin.Context, config *config.Config) {
	var payload services.Payload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := services.TextToText(payload, config.Server.TextClassification, config.Server.Token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arr, err := combineFirstTwoSubArrays(response)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var result string
	for _, element := range arr {
		dict := element.(map[string]interface{})
		label, score := dict["label"].(string), dict["score"].(float64)
		result = fmt.Sprintf("%s%s,%f\n", result, label, score)
	}
	result = strings.TrimSuffix(result, "\n")
	//fmt.Println(result)
	//services.CreatePrompt(c, "TEXT_CLASSIFICAITON", payload.Inputs, result, c.Param("user_id"))

	c.JSON(http.StatusOK, arr)

}

func combineFirstTwoSubArrays(data interface{}) ([]interface{}, error) {
	slice, ok := data.([]interface{})
	if !ok {
		return nil, errors.New("expected data to be a slice of interfaces")
	}

	var result []interface{}

	for _, element := range slice {
		innerSlice, ok := element.([]interface{})
		if !ok {
			return nil, errors.New("unexpected data format in outer slice")
		}

		if len(innerSlice) == 0 {
			continue
		}

		dict := innerSlice[0].(map[string]interface{})
		label, score := dict["label"].(string), dict["score"].(float64)

		result = append(result, map[string]interface{}{"label": label, "score": score})
	}

	return result, nil
}
