package models

import (
	"errors"
	"log"

	"github.com/Arcady1/OZON_Internship_API/utils"
)

func SaveURL(originalUrl string) (string, error) {
	shortURL := utils.GenerateShortURL(originalUrl)

	if utils.GetDataStorageIsDB() == true {
		err := saveURLInDB(shortURL, originalUrl)
		if err != nil {
			return "", err
		}
	} else {
		err := saveURLLocally(shortURL, originalUrl)
		if err != nil {
			return "", err
		}
	}

	return shortURL, nil
}

func saveURLLocally(shortURL string, originalUrl string) error {
	_, exists := allUrls[shortURL]

	if exists == false {
		allUrls[shortURL] = originalUrl
		log.Println("allUrls:", allUrls)
		return nil
	} else {
		return errors.New("Error: the short URL already exists")
	}
}

func saveURLInDB(shortURL string, originalUrl string) error {
	db, err := GetDB()

	if err != nil {
		log.Println(err)
		return errors.New("Error: get DB")
	}

	_, err = db.Query(`
			INSERT INTO urls (short, original)
			VALUES ($1, $2);
		`, shortURL, originalUrl)

	if err != nil {
		log.Println(err)
		return errors.New("Error: insert original URL to the DB by short URL")
	}

	return nil
}
