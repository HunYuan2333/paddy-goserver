package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

type PredictImageData struct {
	ImageId string `json:"imageid"`
	ModelId string `json:"modelid"`
}

func PredictImage(c *gin.Context) {
	var json PredictImageData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Print(err)
		return
	}
	pythonurl := "127.0.0.1/predict_image"
	res, err := http.Post(pythonurl, "application/json", c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error communicating with Python API"})
		log.Printf("Error sending request to Python API: %v", err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// 读取响应体时出错
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Python API"})
		log.Printf("Error reading response from Python API: %v", err)
		return
	}
	c.JSON(http.StatusOK, body)
}
