package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func ShowGrowImage(c *gin.Context) {
	imageid := c.Param("imageId")
	pythonApIUrl := "http://127.0.0.1:5000/showgrowimage/" + imageid
	res, err := http.Get(pythonApIUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Python backend"})
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Python backend"})
		return
	}
	result := gin.H{
		"status": res.StatusCode,
		"body":   string(body),
	}
	contentType := res.Header.Get("mimetype")
	if contentType == "image/jpg" {
		c.Data(res.StatusCode, contentType, body)
	} else {
		c.JSON(res.StatusCode, result)
	}
}
