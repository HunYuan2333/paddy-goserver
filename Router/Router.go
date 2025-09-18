package Router

import (
	"time"

	"paddy-goserver/DataBaseConnection"
	"paddy-goserver/DieaseImage"
	"paddy-goserver/GrowImage"
	"paddy-goserver/MeteorologicalData"
	"paddy-goserver/PredictImage"
	"paddy-goserver/ShowImg"
	"paddy-goserver/UserOperation"
	"paddy-goserver/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
		AllowOrigins:     []string{"http://localhost:8080"},                                                                                // 允许所有源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                                                              // 包括更多HTTP方法
		AllowHeaders:     []string{"Content-Type", "Authorization", "Accept", "Origin", "User-Agent", "Cache-Control", "X-Requested-With"}, // 允许的头部
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // 重要：允许携带cookies
		MaxAge:           12 * time.Hour,
	})

	// 应用 CORS 中间件
	r.Use(corsConfig)

	// 公开路由（不需要鉴权）
	public := r.Group("/")
	{
		public.POST("/userlogin", UserOperation.Userlogin)
		public.POST("/usersignup", UserOperation.UserSignup)
		public.GET("/ShowImg/:ImgType/:imageId", ShowImg.ShowImg)
		public.POST("/GetData/:Type", MeteorologicalData.DataCheck)
	}

	// 需要鉴权的路由
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware()) // 添加鉴权中间件
	{
		protected.POST("/UploadUserImage", UserOperation.UserImageUpload)
		protected.POST("/grow_image/", GrowImage.Upload)
		protected.POST("/DiseaseImage", DieaseImage.DiseaseImageUpload)
		protected.POST("/PredictImage", PredictImage.PredictImage)
		protected.POST("/userlogout", UserOperation.UserLogout) // 退出登录接口
	}

	return r
}
