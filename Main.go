package paddygoserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var database *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:13376035511@tcp(127.0.0.1:3306)/Paddy")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
}
func Main() {
	r := gin.Default()
	r.POST("userlogin", Userlogin)
}
