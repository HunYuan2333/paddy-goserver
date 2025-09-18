package auth

import (
	"errors"
	"strconv"
	"time"

	"paddy-goserver/ConfigInit"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var config, _ = ConfigInit.ReadConfigFile()
var jwtKey = []byte(config.TokenKey)

// Claims JWT声明结构体
type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userID int64) (string, error) {
	// 设置token过期时间为72小时
	expirationTime := time.Now().Add(72 * time.Hour)

	// 创建Claims
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "paddy-goserver",
			Subject:   strconv.FormatInt(userID, 10),
		},
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken 验证JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 检查token是否有效
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// SetTokenCookie 设置JWT token到cookie
func SetTokenCookie(c *gin.Context, token string) {
	// 设置HttpOnly cookie，增强安全性
	c.SetCookie(
		"token",                     // cookie名称
		token,                       // cookie值
		int(72*time.Hour.Seconds()), // 过期时间（秒）
		"/",                         // 路径
		"",                          // 域名（空表示当前域名）
		false,                       // secure（开发环境设为false，生产环境HTTPS时设为true）
		true,                        // httpOnly（防止XSS攻击）
	)
}

// GetTokenFromCookie 从cookie中获取token
func GetTokenFromCookie(c *gin.Context) (string, error) {
	token, err := c.Cookie("token")
	if err != nil {
		return "", err
	}
	return token, nil
}

// GetTokenFromHeader 从 Authorization Header 中获取 token
func GetTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header is missing")
	}

	// Token 通常以 "Bearer " 开头
	const bearerPrefix = "Bearer "
	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		return authHeader[len(bearerPrefix):], nil
	}

	return "", errors.New("invalid authorization header format")
}

// ClearTokenCookie 清除token cookie（用于退出登录）
func ClearTokenCookie(c *gin.Context) {
	c.SetCookie(
		"token",
		"",
		-1, // 设置为过去的时间，立即过期
		"/",
		"",
		false,
		true,
	)
}
