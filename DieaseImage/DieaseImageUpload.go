package DieaseImage

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"paddy-goserver/ConfigInit"
	"path/filepath"
)

func DiseaseImageUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving file from request"})
		return
	}
	config, err := ConfigInit.ReadConfigFile()
	if err != nil {
		log.Print(err)
	}
	filedir := config.FilePath
	filedir = filepath.Join(filedir, "/DiseaseImage")
	err = os.MkdirAll(filedir, 0755)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating directory"})
		return
	}
	myfilepath := filepath.Join(filedir, "/", file.Filename)
	err = c.SaveUploadedFile(file, myfilepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": file.Filename, "message": "上传成功"})
}
