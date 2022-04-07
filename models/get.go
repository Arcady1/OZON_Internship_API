package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Arcady1/OZON_Internship_API/utils"
)

func GetURL(shortURL string) (string, error) {
	var originalUrl string
	var err error

	if utils.GetDataStorageIsDB() == true {
		originalUrl, err = getURLFromDB(db, shortURL)
		if err != nil {
			return "", err
		}
	} else {
		originalUrl, err = getURLLocally(shortURL)
		if err != nil {
			return "", err
		}
	}

	return originalUrl, nil
}

func getURLLocally(shortURL string) (string, error) {
	originalURL, exists := allUrls[shortURL]

	if exists == false {
		return "", errors.New("Error: the short URL doesn't exist")
	} else {
		return originalURL, nil
	}
}

func getURLFromDB(db *sql.DB, shortURL string) (string, error) {
	data, err := db.Query(`
			SELECT original FROM urls
			WHERE short = $1;
		`, shortURL)

	var originalURL string

	for data.Next() {
		err = data.Scan(&originalURL)
		if err != nil {
			fmt.Println(err)
			return "", errors.New("Error: scanning results after getting the original URL form the DB by short URL")
		}
	}

	if err != nil {
		fmt.Println(err)
		return "", errors.New("Error: get original URL form the DB by short URL")
	}

	return originalURL, nil
}
