package PredictImage

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"paddy-goserver/ConfigInit"
)

func ShowPredictGrowImage(c *gin.Context) {
	Imgid := c.Param("imgid")
	pythonurl := "http://127.0.0.1:5000/show_predict_grow_image/" + Imgid
	res, err := http.Get(pythonurl)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "python error",
		})
		log.Print(err)
		return
	}
	config, err := ConfigInit.ReadConfigFile()
	if err != nil {
		log.Print(err)
	}
	filedir := config.FilePath
	filedir = filedir + "/" + "PredictGrowImage"
	err = os.MkdirAll(filedir, 0755)
	if err != nil {
		log.Print(err)
	}
	myfilepath := filedir + "/" + Imgid + ".jpg"
	body, err := ioutil.ReadAll(res.Body)
	f, err := os.Create(myfilepath)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)
	_, err = f.Write(body)
	result := gin.H{
		"status": res.StatusCode,
		"body":   string(body),
	}
	ContentType := res.Header.Get("mimetype")
	if ContentType == "image/jpg" {
		c.Data(res.StatusCode, ContentType, body)
	} else {
		c.JSON(res.StatusCode, result)
	}
}
