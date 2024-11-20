package GrowImage

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error retrieving file from request"})
		return
	}
	filedir := os.Getenv("PADDY_SERVER_FILE_PATH")
	filedir = filepath.Join(filedir, "/GrowImage")
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
