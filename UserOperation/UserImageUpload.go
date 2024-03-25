package UserOperation

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
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
	filedir = filedir + "/" + "UserImage"
	err = os.MkdirAll(filedir, 0755)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	imgurl := filedir + "/" + userid + ".jpg"
	dst, err := os.Create(imgurl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer dst.Close()
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open uploaded file"})
		return
	}
	defer src.Close()

	if _, err := io.Copy(dst, src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to copy file contents"})
		return
	}
	stmt, err := database.Prepare("UPDATE User SET imgurl=? WHERE Userid=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(imgurl, userid)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
