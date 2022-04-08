package models

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectToDB() (*sql.DB, error) {
	USER := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	DBNAME := os.Getenv("DB_NAME")
	SSLMODE := os.Getenv("DB_SSLMODE")

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", USER, PASSWORD, DBNAME, SSLMODE)
	database, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return nil, errors.New("Error: connect to the Database")
	} else {
		db = database
	}

	return db, nil
}

func Migrate() error {
	err := dropTable()
	if err != nil {
		return errors.New("Error: drop table on migration")
	}

	err = createTable()
	if err != nil {
		return errors.New("Error: create table on migration")
	}

	return nil
}

func createTable() error {
	fmt.Println("HERE_2")
	_, err := db.Exec(`
			CREATE TABLE urls (
			short varchar(11) NOT NULL,
			original varchar(200),
			PRIMARY KEY (short));
		`)

	return err
}

func dropTable() error {
	fmt.Println("HERE_1")
	_, err := db.Exec(`DROP TABLE IF EXISTS urls;`)
	fmt.Println("HERE_1", err)
	return err
}
