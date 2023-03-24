package repository

import (
	"database/sql"
	"fmt"
	"log"
)

var Db *sql.DB

func ConnectDatabase() *sql.DB {
	var err error
	Db, err = sql.Open("mysql", "user:password@tcp(mysql:3306)/db"+"?charset=utf8&parseTime=true")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	fmt.Println("Database connection succeded!!")
	return Db
}
