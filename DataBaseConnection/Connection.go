package DataBaseConnection

import (
	"fmt"
	"paddy-goserver/ConfigInit"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var config *ConfigInit.Config

func init() {
	config, _ = ConfigInit.ReadConfigFile()
}
func SetupDatabase() error {
	db, err := sqlx.Open(config.DriverName, config.DriverCommand)
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
