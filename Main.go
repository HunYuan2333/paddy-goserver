package paddygoserver

import (
	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Main() {
	r := gin.Default()
	r.POST("userlogin", userlogin)
}
