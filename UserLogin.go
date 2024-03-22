package main

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Userlogin(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prepstmt := "SELECT COUNT(*) FROM User WHERE Username =? AND Password = ?"
	stmt, preperr := database.Prepare(prepstmt)
	if preperr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error preparing statement"})
		return
	}
	var count int64
	// 使用 _ 作为占位符，因为我们只关心查询是否成功
	err := stmt.QueryRow(json.Username, json.Password).Scan(&count)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	case err != nil:
		log.Printf("Unexpected error while fetching user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user from the database"})
	default:
		if count > 0 {
			c.JSON(http.StatusOK, gin.H{"status": "200"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		}
	}
}
