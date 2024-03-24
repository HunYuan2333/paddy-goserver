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
	r.GET("/ShowGrowImage/<imageId>", GrowImage.ShowGrowImage)
	r.POST("/DiseaseImage", DieaseImage.DiseaseImageUpload)
	r.GET("/ShowDiseaseImage/<imageId>", PredictImage.ShowPredictDiseaseImage)
	r.POST("/PredictImage", PredictImage.PredictImage)
	r.GET("/ShowPredictGrowImage/<imageId>", PredictImage.ShowPredictGrowImage)
	r.GET("/ShowPredictDiseaseImage/<imageId>", PredictImage.ShowPredictDiseaseImage)
	return r
}
