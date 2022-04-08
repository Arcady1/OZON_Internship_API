package models

import (
	"database/sql"
	"errors"
	"log"

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
	log.Println("Getting data locally")

	originalURL, exists := allUrls[shortURL]

	if exists == false {
		return "", errors.New("Error: the short URL doesn't exist")
	} else {
		return originalURL, nil
	}
}

func getURLFromDB(db *sql.DB, shortURL string) (string, error) {
	log.Println("Getting data from DB")

	var originalURL string
	var isData bool = false

	data, err := db.Query(`
			SELECT original FROM urls
			WHERE short = $1;
		`, shortURL)

	for data.Next() {
		isData = true
		err = data.Scan(&originalURL)
		if err != nil {
			log.Println(err)
			return "", errors.New("Error: scanning results after getting the original URL form the DB by short URL")
		}
	}

	if isData == false {
		return "", errors.New("Error: the short URL doesn't exist")
	}

	if err != nil {
		log.Println(err)
		return "", errors.New("Error: get original URL form the DB by short URL")
	}

	return originalURL, nil
}
