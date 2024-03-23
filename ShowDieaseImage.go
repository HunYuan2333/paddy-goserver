package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func ShowDiseaseImage(c *gin.Context) {
	imgid := c.Param("imageId")
	pythonurl := "http://127.0.0.1:5000/show_disease_image/" + imgid
	res, err := http.Get(pythonurl)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "连接python 服务器出错",
		})
		log.Print(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from Python backend"})
		log.Print(err)
		return
	}
	result := gin.H{
		"status": res.StatusCode,
		"body":   string(body),
	}
	content_type := res.Header.Get("mimetype")
	if content_type == "image/jpg" {
		c.Data(res.StatusCode, content_type, body)
	} else {
		c.JSON(res.StatusCode, result)
	}
}
