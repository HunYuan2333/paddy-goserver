package GrowImage

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func ShowGrowImage(c *gin.Context) {
	imgid := c.Param("imageId") + ".jpg"
	Imgdir := os.Getenv("PADDY_SERVER_FILE_PATH")
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
	content_type := "image/jpg"
	c.Data(http.StatusOK, content_type, body)
}
