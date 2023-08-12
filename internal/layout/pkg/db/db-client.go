package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type DbClient struct{
	db *sql.DB
}
func GetDbInstance() *DbClient{
    instance:=new(DbClient)
	instance.init()
	return instance
}

func(d *DbClient) init(){
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "golang_api",
		AllowNativePasswords: true,
	}

	// Get a database handle.
	var err error
	d.db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := d.db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}