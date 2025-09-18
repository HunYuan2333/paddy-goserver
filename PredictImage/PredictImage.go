package PredictImage

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	ImageId string `json:"imageid"`
	ModelId string `json:"modelid"`
}

func PredictImage(c *gin.Context) {
	// Print the raw request body before binding
	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error reading raw request body: %v", err)
		// Continue processing even if reading raw body fails, as ShouldBindJSON might still work
	} else {
		log.Printf("Raw Request Body: %s", string(bodyBytes))
		// Restore the body so ShouldBindJSON can read it
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	var jsondata Data
	if err := c.ShouldBindJSON(&jsondata); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		log.Print(err)
		return
	}

	pythonurl := "http://127.0.0.1:5050/predict_image"
	// Use the read bodyBytes in the new request
	res, err := http.Post(pythonurl, "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error communicating with Python API"})
		log.Printf("Error sending request to Python API: %v", err)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)
	responseBodyBytes, responseErr := ioutil.ReadAll(res.Body) // Changed variable names
	if responseErr != nil {                                    // Changed variable name
		// 读取响应体时出错
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Python API"})
		log.Printf("Error reading response from Python API: %v", responseErr) // Changed variable name
		return
	}
	log.Printf("Response from Python API: %s", string(responseBodyBytes))

	// Unmarshal the byte slice into a map
	var result map[string]interface{}
	if err := json.Unmarshal(responseBodyBytes, &result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response from Python API"})
		log.Printf("Error unmarshalling response from Python API: %v", err)
		return
	}

	c.JSON(http.StatusOK, result) // Send the parsed map as JSON
}
