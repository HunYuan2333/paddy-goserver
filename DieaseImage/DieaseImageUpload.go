package DieaseImage

import (
	"log"
	"net/http"
	"os"
	"paddy-goserver/ConfigInit"
	"path/filepath"

	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading config file"})
		return
	}
	filedir := filepath.Join(config.FilePath, "disease")
	err = os.MkdirAll(filedir, 0755)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error creating directory"})
		return
	}
	myfilepath := filepath.Join(filedir, file.Filename)
	err = c.SaveUploadedFile(file, myfilepath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving file"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": file.Filename, "message": "上传成功"})
}
