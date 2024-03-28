package UserOperation

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"paddy-goserver/DataBaseConnection"
)

// Login 结构体用于定义用户登录信息
type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

var database *sqlx.DB

func init() {
	if err := DataBaseConnection.SetupDatabase(); err != nil {
		panic(err)
	}
	database = DataBaseConnection.GetDatabase()
}

// Userlogin 处理用户登录请求
// c *gin.Context: Gin框架的上下文对象，用于处理HTTP请求和响应
func Userlogin(c *gin.Context) {
	var json Login // 用于存储从请求中解析的JSON登录信息

	// 尝试从请求体中解析JSON格式的登录信息
	if err := c.ShouldBindJSON(&json); err != nil {
		// 如果解析失败，返回状态码400（Bad Request）和错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 准备数据库查询语句
	prepstmt := "SELECT COUNT(*),Userid,imgurl FROM User WHERE Username =? AND Password = ? GROUP BY Userid, imgurl"
	stmt, preperr := database.Prepare(prepstmt)
	if preperr != nil {
		// 如果准备语句失败，返回状态码500（Internal Server Error）和错误信息
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error preparing statement"})
		log.Print(preperr)
		return
	}
	var count int64 // 用于存储查询结果的计数
	var Userid int64
	var imgurl string
	// 执行查询，检查用户名和密码是否匹配
	err := stmt.QueryRow(json.Username, json.Password).Scan(&count, &Userid, &imgurl)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		// 如果查询无结果，返回状态码401（Unauthorized）和错误信息
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	case err != nil:
		// 如果查询过程中发生其他错误，记录错误日志并返回状态码500和错误信息
		log.Printf("Unexpected error while fetching user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user from the database"})
	default:
		// 根据查询结果计数判断登录是否成功，并返回相应的状态码和信息
		if count > 0 {
			c.JSON(http.StatusOK, gin.H{"code": "200",
				"id":       Userid,
				"imgurl":   imgurl,
				"username": Login{Username: json.Username}})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	}
}
