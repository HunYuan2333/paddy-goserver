package main

import (
	"log"
	"net/http"
	"paddy-goserver/Router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
	err = r.Run(":8082") // 修改端口为8082，避免与Python服务器(5000)冲突
	if err != nil {
		log.Print(err)
		return
	}
}
