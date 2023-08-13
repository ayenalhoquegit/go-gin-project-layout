package db

import (
	"database/sql"
	"fmt"
	"log"

	//"github.com/ayenalhoquegit/go-gin-project-layout/pkg/config"
	"github.com/go-sql-driver/mysql"
)

type DbClient struct {
	DB *sql.DB
}

func GetDbInstance() *DbClient {
	instance := new(DbClient)
	instance.init()
	return instance
}
const db_user="root"
func (d *DbClient) init() {
	cfg := mysql.Config{
		User:                 db_user,
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "golang_api",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var err error
	d.DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := d.DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
