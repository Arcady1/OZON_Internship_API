package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	// TODO
	USER := "arch"
	PASSWORD := "1029"
	DBNAME := "ozon"
	SSLMODE := "disable"

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", USER, PASSWORD, DBNAME, SSLMODE)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println(err)
		DB = nil
	} else {
		DB = db
	}
}
