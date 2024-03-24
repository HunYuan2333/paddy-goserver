package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"paddy-goserver/Router"
)

func main() {
	var err error
	if err != nil {
		log.Print(err)
		panic(err)
	}
	r := Router.InitRouter()
	err = r.Run(":8080")
	if err != nil {
		log.Print(err)
		return
	}
}
