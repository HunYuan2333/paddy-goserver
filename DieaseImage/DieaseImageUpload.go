package DieaseImage

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"io/ioutil"
	"log"
	"net/http"
)

func DiseaseImageUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err == nil {
		if file != nil {
			imgbyte, err := file.Open()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "读取上传文件失败"})
				log.Print(err)
				return
			}
			defer imgbyte.Close()
			res, err := http.Post("127.0.0.1:8080/upload_disease_image", "multipart/form-data", imgbyte)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "转发图片到Python服务失败"})
				log.Print(err)
				return
			}
			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)
			responseJSON := make(map[string]interface{})
			json.Unmarshal(body, &responseJSON)
			c.JSON(http.StatusOK, responseJSON)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "未接收到文件"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "获取上传文件失败"})
		return
	}
}
