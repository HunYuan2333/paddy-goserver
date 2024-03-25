package Router

import (
	"github.com/gin-gonic/gin"
	"paddy-goserver/DataBaseConnection"
	"paddy-goserver/DieaseImage"
	"paddy-goserver/GrowImage"
	"paddy-goserver/PredictImage"
	"paddy-goserver/UserOperation"
)

func init() {
	if err := DataBaseConnection.SetupDatabase(); err != nil {
		panic(err)
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/userlogin", UserOperation.Userlogin)
	r.POST("/usersignup", UserOperation.UserSignup)
	r.POST("/grow_image/", GrowImage.GrowImageUpload)
	r.GET("/ShowImg/<ImgType>/<imageId>", GrowImage.ShowGrowImage)
	r.POST("/DiseaseImage", DieaseImage.DiseaseImageUpload)
	r.POST("/PredictImage", PredictImage.PredictImage)
	return r
}
