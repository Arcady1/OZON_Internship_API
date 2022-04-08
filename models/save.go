package models

import (
	"errors"
	"log"

	"github.com/Arcady1/OZON_Internship_API/utils"
)

func SaveOriginalURL(originalUrl string) (string, error) {
	var (
		shortURL string
		err      error
	)

	// TODO
	if utils.GetDataStorageIsDB() == true {
		err = saveURLInDB(shortURL, originalUrl)
		if err != nil {
			return "", err
		}
	} else {
		shortURL, err = saveOriginalURLLocally(originalUrl)

		shortURL, err = saveOriginalURLLocally(originalUrl)
		if err != nil {
			return "", err
		}
	}

	return shortURL, nil
}

func saveOriginalURLLocally(originalUrl string) (string, error) {
	shortURL := ifOriginalURLAlreadySavedLocally(originalUrl)

	// If originalUrl already saved locally, return existing shortURL
	if shortURL != "" {
		return shortURL, nil
	}

	// If originalUrl is not saved, generate new shortURL
	shortURL = utils.GenerateShortURL(originalUrl)
	exists := ifShortURLAlreadySavedLocally(shortURL)

	// If shortURL doesn't not exist
	if exists == false {
		allUrls[shortURL] = originalUrl
		return shortURL, nil
	} else {
		return "", errors.New("Error: the generated short URL already exists")
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
