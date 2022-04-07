package models

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB = nil
var dbError error = nil

func init() {
	USER := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_DBNAME")
	SSLMODE := os.Getenv("DB_SSLMODE")

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", USER, PASSWORD, DBNAME, SSLMODE)
	database, err := sql.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println(err)
		dbError = errors.New("Error: connect to the Database")
		db = nil
		return
	} else {
		db = database
	}

	err = dropTable()
	if err != nil {
		fmt.Println(err)
		dbError = errors.New("Error: drop table on migration")
		db = nil
		return
	}

	err = createTable()
	if err != nil {
		fmt.Println(err)
		dbError = errors.New("Error: create table on migration")
		db = nil
		return
	}
}

func GetDB() (*sql.DB, error) {
	return db, dbError
}

func createTable() error {
	_, err := db.Exec(`
			CREATE TABLE urls (
			short varchar(11) NOT NULL,
			original varchar(200),
			PRIMARY KEY (short));
		`)

	return err
}

func dropTable() error {
	_, err := db.Exec(`DROP TABLE IF EXISTS urls;`)

	return err
}
