package paddygoserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func userlogin(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO:SQl operate
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
