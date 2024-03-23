package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io/ioutil"
	"log"
	"net/http"
)

func GrowImageUpload(c *gin.Context) {
	if c.Request.Method == http.MethodPost {
		file, err := c.FormFile("file")
		if err == nil {
			if file != nil {
				// 读取文件内容
				imgBytes, err := file.Open()
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "读取上传文件失败"})
					log.Print(err)
					return
				}
				defer imgBytes.Close()
				// 将图片内容发送到Python Flask应用
				resp, err := http.Post(
					"http://localhost:5000/upload_grow_image",
					file.Header.Get("Content-Type"),
					imgBytes,
				)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "转发图片到Python服务失败"})
					log.Print(err)
					return
				}
				defer resp.Body.Close()

				body, _ := ioutil.ReadAll(resp.Body)
				responseJSON := make(map[string]interface{})
				json.Unmarshal(body, &responseJSON)
				c.JSON(http.StatusOK, responseJSON)
				return
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"message": "未接收到文件"})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "获取上传文件失败"})
			return
		}
	}

	c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "仅支持POST方法"})
}
