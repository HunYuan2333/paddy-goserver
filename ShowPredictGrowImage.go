package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func ShowPredictGrowImage(c *gin.Context) {
	Imgid := c.Param("imgid")
	pythonurl := "http://127.0.0.1:5000/show_predict_grow_image/" + Imgid
	res, err := http.Get(pythonurl)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "python error",
		})
		log.Print(err)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	result := gin.H{
		"status": res.StatusCode,
		"body":   string(body),
	}
	ContentType := res.Header.Get("mimetype")
	if ContentType == "image/jpg" {
		c.Data(res.StatusCode, ContentType, body)
	} else {
		c.JSON(res.StatusCode, result)
	}
}
