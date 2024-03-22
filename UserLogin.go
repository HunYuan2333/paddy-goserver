package paddygoserver

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
	stmt := "SELECT * FROM User WHERE username =? AND password = ?"
	row := database.QueryRow(stmt, json.Username, json.Password)
	var DB Login
	err := row.Scan(&DB.Username, &DB.Password)
	switch {
	case err == sql.ErrNoRows:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Interstitials"})
	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user from the database"})
	default:
		c.JSON(http.StatusOK, gin.H{"status": "200"})

	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
