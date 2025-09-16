package middlewares

import (
	"net/http"
	"paddy-goserver/auth"

	"github.com/gin-gonic/gin"
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

// AuthMiddleware JWT鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从cookie中获取token
		tokenString, err := auth.GetTokenFromCookie(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "未找到认证token",
				"code":  401,
			})
			c.Abort()
			return
		}

		// 验证token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无效的认证token",
				"code":  401,
			})
			c.Abort()
			return
		}

		// 将用户ID存储到上下文中，供后续处理函数使用
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

// OptionalAuthMiddleware 可选的鉴权中间件（用于某些接口可以选择性验证身份）
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 尝试从cookie获取token
		tokenString, err := auth.GetTokenFromCookie(c)
		if err == nil {
			// 如果存在token，尝试验证
			if claims, err := auth.ValidateToken(tokenString); err == nil {
				c.Set("user_id", claims.UserID)
				c.Set("authenticated", true)
			}
		}
		// 无论是否有有效token都继续执行
		c.Next()
	}
}
