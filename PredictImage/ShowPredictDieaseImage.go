package PredictImage

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ShowPredictDiseaseImage(c *gin.Context) {
	Imgid := c.Param("imgid")
	pythonurl := "http://127.0.0.1:5000/show_predict_disease_image/" + Imgid
	res, err := http.Get(pythonurl)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusBadRequest,
			"message": "python error",
		})
		log.Print(err)
		return
	}
	filedir := os.Getenv("PADDY_SERVER_FILE_PATH")
	filedir = filedir + "/" + "PredictDiseaseImage"
	err = os.MkdirAll(filedir, 0755)
	if err != nil {
		log.Print(err)
	}
	myfilepath := filedir + "/" + Imgid + ".jpg"
	body, err := ioutil.ReadAll(res.Body)
	f, err := os.Create(myfilepath)
	defer f.Close()
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
