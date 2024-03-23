package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

var database *sqlx.DB

func setupDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", "root:13376035511@tcp(127.0.0.1:3306)/Paddy")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return nil, err
	}
	// 可以添加Ping或者其他的健康检查，确保数据库连接可用
	if err = db.Ping(); err != nil {
		fmt.Println("ping mysql failed,", err)
		return nil, err
	}
	return db, nil
}

func initRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()
	r.POST("/userlogin", Userlogin)
	r.POST("/usersignup", UserSignUp)
	r.POST("/grow_image/", GrowImageUpload)
	r.GET("/ShowGrowImage/<imageId>", ShowGrowImage)
	r.POST("/DiseaseImage", DiseaseImageUpload)
	r.GET("/ShowDiseaseImage/<imageId>", ShowDiseaseImage)
	r.POST("/PredictImage", PredictImage)
	r.GET("/ShowPredictGrowImage/<imageId>", ShowPredictGrowImage)
	r.GET("/ShowPredictDiseaseImage/<imageId>", ShowPredictDiseaseImage)
	return r
}
func main() {
	var err error
	database, err = setupDatabase()
	if err != nil {
		log.Print(err)
		panic(err)
	}
	r := initRouter(database)
	err = r.Run(":8080")
	if err != nil {
		log.Print(err)
		return
	}
}
