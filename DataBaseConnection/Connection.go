package DataBaseConnection

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func SetupDatabase() error {
	db, err := sqlx.Open("mysql", "root:1234abcd@tcp(127.0.0.1:3306)/paddy")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return err
	}

	if err = db.Ping(); err != nil {
		fmt.Println("ping mysql failed,", err)
		return err
	}
	DB = db
	return nil
}

func GetDatabase() *sqlx.DB {
	return DB
}
