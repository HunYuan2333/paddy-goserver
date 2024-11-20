package Router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"paddy-goserver/DataBaseConnection"
	"paddy-goserver/DieaseImage"
	"paddy-goserver/GrowImage"
	"paddy-goserver/MeteorologicalData"
	"paddy-goserver/PredictImage"
	"paddy-goserver/ShowImg"
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

	// 配置 CORS
	corsConfig := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},                                                                                                   // 允许的源
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},                                                                                                  // 包括 OPTIONS 方法
		AllowHeaders:     []string{"Content-Type", "Authorization", "Accept", "Referer", "User-Agent", "Sec-Ch-Ua", "Sec-Ch-Ua-Mobile", "Sec-Ch-Ua-Platform"}, // 允许的头部
		ExposeHeaders:    []string{},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})

	// 应用 CORS 中间件
	r.Use(corsConfig)

	// 定义路由
	r.POST("/UploadUserImage", UserOperation.UserImageUpload)
	r.POST("/userlogin", UserOperation.Userlogin)
	r.POST("/usersignup", UserOperation.UserSignup)
	r.POST("/grow_image/", GrowImage.Upload)
	r.GET("/ShowImg/:ImgType/:imageId", ShowImg.ShowImg)
	r.POST("/DiseaseImage", DieaseImage.DiseaseImageUpload)
	r.POST("/PredictImage", PredictImage.PredictImage)
	r.POST("/GetData/:Type", MeteorologicalData.DataCheck)

	return r
}
