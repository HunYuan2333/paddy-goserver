package GrowImage

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"paddy-goserver/ConfigInit"
	"path/filepath"
)

func ShowGrowImage(c *gin.Context) {
	imgid := c.Param("imageId") + ".jpg"
	config, err := ConfigInit.ReadConfigFile()
	if err != nil {
		log.Print(err)
	}
	Imgdir := config.FilePath
	ImgPath := filepath.Join(Imgdir, "/GrowImage")
	if err := os.MkdirAll(ImgPath, 0755); err != nil {
		// 处理错误
		log.Fatal(err)
	}
	ImgPath = filepath.Join(ImgPath, "/", imgid)
	body, err := ioutil.ReadFile(ImgPath)
	if err != nil {
		c.JSON(404, gin.H{
			"status": 404,
			"body":   "图片不存在",
		})
		return
	}
	contentType := "image/jpg"
	c.Data(http.StatusOK, contentType, body)
}
