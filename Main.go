package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"paddy-goserver/Router"
)

func main() {
	var err error
	if err != nil {
		log.Print(err)
		panic(err)
	}
	r := Router.InitRouter()
	// 设置跨域中间件

	// 示例路由
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	// 启动服务器
	err = r.Run(":5000")
	if err != nil {
		log.Print(err)
		return
	}
}
