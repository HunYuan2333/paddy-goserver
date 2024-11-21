package UserOperation

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"paddy-goserver/ConfigInit"
	"paddy-goserver/DataBaseConnection"
	"strconv"
	"time"
)

// Login 结构体用于定义用户登录信息
type Login struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

var config, _ = ConfigInit.ReadConfigFile()
var database *sqlx.DB
var jwtKey = []byte(config.TokenKey)

func init() {
	if err := DataBaseConnection.SetupDatabase(); err != nil {
		panic(err)
	}
	database = DataBaseConnection.GetDatabase()
}
func setTokenCookie(c *gin.Context, token string) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true, // 如果使用 HTTPS，设置为 true
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(c.Writer, cookie)
}
func generateToken(userId int64) (string, error) {
	// 创建一个我们自己的声明数据结构
	expiresAt := jwt.NewNumericDate(time.Now().Add(time.Hour * 72))
	claims := &jwt.RegisteredClaims{
		Issuer:    "Login",
		Subject:   strconv.FormatInt(userId, 10),
		ExpiresAt: expiresAt, // 令牌有效期为72小时
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用指定的密钥签名并获得完整的编码后的字符串
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatalf("Failed to sign token: %v", err)
		return "", err
	}

	return signedToken, nil
}

func Userlogin(c *gin.Context) {
	var json Login // 用于存储从请求中解析的JSON登录信息

	// 尝试从请求体中解析JSON格式的登录信息
	if err := c.ShouldBindJSON(&json); err != nil {
		// 如果解析失败，返回状态码400（Bad Request）和错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 准备数据库查询语句
	prepstmt := "SELECT Userid,imgurl,Password FROM User WHERE Username =? GROUP BY Userid, imgurl,Password"
	stmt, preperr := database.Prepare(prepstmt)
	if preperr != nil {
		// 如果准备语句失败，返回状态码500（Internal Server Error）和错误信息
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error preparing statement"})
		log.Print(preperr)
		return
	} // 用于存储查询结果的计数
	var Userid int64
	var imgurl string
	var hashedPassword []byte
	// 执行查询，检查用户名和密码是否匹配
	err := stmt.QueryRow(json.Username).Scan(&Userid, &imgurl, &hashedPassword)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		// 如果查询无结果，返回状态码401（Unauthorized）和错误信息
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	case err != nil:
		// 如果查询过程中发生其他错误，记录错误日志并返回状态码500和错误信息
		log.Printf("Unexpected error while fetching user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user from the database"})
	default:
		hasherr := bcrypt.CompareHashAndPassword(hashedPassword, []byte(json.Password))
		// 根据查询结果计数判断登录是否成功，并返回相应的状态码和信息
		if hasherr == nil {
			var token, _ = generateToken(Userid)
			setTokenCookie(c, token)
			c.JSON(http.StatusOK, gin.H{"code": "200",
				"id":       Userid,
				"imgurl":   imgurl,
				"username": json.Username,
				"token":    token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	}
}
