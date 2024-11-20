package UserOperation

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func UserImageUpload(c *gin.Context) {
	// 获取表单数据
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	files := form.File["file"]
	if len(files) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file attached"})
		return
	}
	file := files[0]
	userid := form.Value["userid"][0]
	filedir := os.Getenv("PADDY_SERVER_FILE_PATH")
	log.Print(filedir)
	filedir = filedir + "/" + "UserImage"
	err = os.MkdirAll(filedir, 0755)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Print(err)
		return
	}
	imgurl := filedir + "/" + userid + ".jpg"
	dst, err := os.Create(imgurl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
		}
	}(dst)
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open uploaded file"})
		return
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
		}
	}(src)

	if _, err := io.Copy(dst, src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file contents"})
		return
	}
	stmt, err := database.Prepare("UPDATE User SET imgurl=? WHERE Userid=?")
	if err != nil {
		log.Fatal(err)
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
		}
	}(stmt)
	databaseimgurl := userid + ".jpg"
	_, err = stmt.Exec(databaseimgurl, userid)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
