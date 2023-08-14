package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ayenalhoquegit/go-gin-project-layout/pkg/config"
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

//const db_user = "root"

func (d *DbClient) init() {
	cfg := mysql.Config{
		User:                 config.GetEnvValue("DB_USER"),
		Passwd:               config.GetEnvValue("DB_PASSWORD"),
		Net:                  "tcp",
		Addr:                 config.GetEnvValue("DB_HOST"),
		DBName:               config.GetEnvValue("DB_NAME"),
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
