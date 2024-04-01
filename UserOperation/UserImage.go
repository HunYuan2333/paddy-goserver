package UserOperation

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func UserImage(c *gin.Context) {
	Imgurl := c.Param("imageId")
	filename := os.Getenv("PADDY_SERVER_FILE_PATH") + "/" + "UserImage" + "/" + Imgurl
	imgData, err := ioutil.ReadFile(filename)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	// 根据图片扩展名设置正确的MIME类型
	ext := filepath.Ext(filename)
	mimeType := ""
	switch ext {
	case ".jpg", ".jpeg":
		mimeType = "image/jpeg"
	case ".png":
		mimeType = "image/png"
	case ".gif":
		mimeType = "image/gif"
	default:
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Data(http.StatusOK, mimeType, imgData)

	/*
		prepstmt := "SELECT COUNT(*) FROM User WHERE imgurl=?"
		stmt, preperr := database.Prepare(prepstmt)
		if preperr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error preparing statement"})
			log.Print(preperr)
			return
		}
		var count int64
		err := stmt.QueryRow(Imgurl).Scan(&count)
		switch {
		case errors.Is(err, sql.ErrNoRows):
			// 如果查询无结果，返回状态码401（Unauthorized）和错误信息
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		case err != nil:
			// 如果查询过程中发生其他错误，记录错误日志并返回状态码500和错误信息
			log.Printf("Unexpected error while fetching user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user from the database"})
		default:
			// 根据查询结果计数判断登录是否成功，并返回相应的状态码和信息
			if count > 0 {
				Imgurl := fmt.Sprintf("%s.png", Imgurl)
				ImgDir := os.Getenv("PADDY_SERVER_FILE_PATH")
				FileImgUrl := filepath.Join(ImgDir, fmt.Sprintf("UserImage/%s", Imgurl))
				mgData, err := ioutil.ReadFile(FileImgUrl)
				if err != nil {
					c.String(http.StatusInternalServerError, "Error reading image file: %v", err)
					return
				}
				c.Data(http.StatusOK, "image/png", mgData)
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			}

		}

	*/
}
