package main

import (
	"github.com/gin-contrib/cors"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"paddy-goserver/Router"
)

func main() {
	var err error
	if err != nil {
		log.Print(err)
		panic(err)
	}
	r := Router.InitRouter()
	// 创建CORS配置对象
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}                 // 允许特定源
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"} // 允许的方法
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}         // 允许的头部
	corsConfig.AllowCredentials = true                                          // 允许带凭证的请求（如 cookies）

	r.Use(cors.New(corsConfig))
	err = r.Run(":5000")
	if err != nil {
		log.Print(err)
		return
	}
}
