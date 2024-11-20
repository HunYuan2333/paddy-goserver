package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//goland:noinspection GoUnusedExportedFunction
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")

		if origin != "" {
			// 允许任何来源进行跨域访问
			c.Header("Access-Control-Allow-Origin", "*")

			// 允许所有HTTP方法
			c.Header("Access-Control-Allow-Methods", "*")

			// 允许所有请求头
			c.Header("Access-Control-Allow-Headers", "*")

			// 暴露其他必要的响应头
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")

			// 允许携带凭据（cookies）
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		// 对于预检请求（OPTIONS方法），直接返回204状态码
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			// 继续处理非预检请求
			c.Next()
		}
	}
}
