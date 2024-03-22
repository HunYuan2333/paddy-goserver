package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SignUpData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserSignUp(c *gin.Context) {
	var json SignUpData
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prepstmt := "INSERT INTO User(Username,Password,imgurl) VALUES (?,?,?)"
	stmt, errpre := database.Prepare(prepstmt)
	if errpre != nil {
		fmt.Print("error in datacommand prepare")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register customer"})
		return
	}
	_, err := stmt.Exec(json.Username, json.Password, nil)
	if err != nil { // 处理错误...
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register customer"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Successfully registered customer"})
}
