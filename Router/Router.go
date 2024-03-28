package Router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"paddy-goserver/DataBaseConnection"
	"paddy-goserver/DieaseImage"
	"paddy-goserver/GrowImage"
	"paddy-goserver/PredictImage"
	"paddy-goserver/UserOperation"
	"time"
)

func init() {
	if err := DataBaseConnection.SetupDatabase(); err != nil {
		panic(err)
	}
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许任何域名访问，生产环境应替换为具体的域名
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
	r.Use(corsConfig)
	r.POST("/UploadUserImage", UserOperation.UserImageUpload)
	r.POST("/userlogin", UserOperation.Userlogin)
	r.POST("/usersignup", UserOperation.UserSignup)
	r.POST("/grow_image/", GrowImage.GrowImageUpload)
	r.GET("/ShowImg/<ImgType>/<imageId>", GrowImage.ShowGrowImage)
	r.POST("/DiseaseImage", DieaseImage.DiseaseImageUpload)
	r.POST("/PredictImage", PredictImage.PredictImage)
	return r
}
