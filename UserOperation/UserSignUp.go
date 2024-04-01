package UserOperation

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type SignUpData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HashPassword(password string) (string, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedpassword), nil
}
func UserSignup(c *gin.Context) {
	var json SignUpData

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prepStmtCheck := "SELECT COUNT(*) FROM User WHERE Username = ?"
	row := database.QueryRow(prepStmtCheck, json.Username)
	var count int
	if err := row.Scan(&count); err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("Error checking for duplicate username: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
		return
	}
	if count > 0 {
		// 用户名已存在，返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	HashedPassword, hasherr := HashPassword(json.Password)
	if hasherr != nil {
		log.Printf("Error hashing password: %v", hasherr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process request"})
		return
	}
	prepstmt := "INSERT INTO User(Username,Password,imgurl) VALUES (?,?,?)"
	stmt, errpre := database.Prepare(prepstmt)
	if errpre != nil {
		fmt.Print("error in datacommand prepare")
		log.Print(errpre)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register customer"})
		return
	}
	_, err := stmt.Exec(json.Username, HashedPassword, "D:\\work_space\\go\\paddy-goserver\\Image\\no.png")
	if err != nil { // 处理错误...
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register customer"})
		return
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "注册成功",
	})
}
