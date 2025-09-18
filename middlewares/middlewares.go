package middlewares

import (
	"log"
	"net/http"
	"paddy-goserver/auth"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware JWT鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Let preflight OPTIONS requests pass through so CORS middleware can set headers
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
		log.Println("--- AuthMiddleware: New request received ---")
		// 从Header中获取token
		tokenString, err := auth.GetTokenFromHeader(c)
		if err != nil {
			log.Printf("AuthMiddleware: Error getting token from header: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "未找到认证token",
				"code":  401,
			})
			c.Abort()
			return
		}
		log.Printf("AuthMiddleware: Token received: %s\n", tokenString)

		// 验证token
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			log.Printf("AuthMiddleware: Error validating token: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "无效的认证token",
				"code":  401,
			})
			c.Abort()
			return
		}

		log.Printf("AuthMiddleware: Token validated successfully for user_id: %d\n", claims.UserID)
		// 将用户ID存储到上下文中，供后续处理函数使用
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}

// OptionalAuthMiddleware 可选的鉴权中间件（用于某些接口可以选择性验证身份）
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Let preflight OPTIONS requests pass through
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
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
