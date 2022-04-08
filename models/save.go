package models

import (
	"errors"

	"github.com/Arcady1/OZON_Internship_API/utils"
)

func SaveOriginalURL(originalUrl string) (string, error) {
	var (
		shortURL string
		err      error
	)

	if utils.GetDataStorageIsDB() == true {
		shortURL, err = saveOriginalURLInDB(originalUrl)
		if err != nil {
			return "", err
		}
	} else {
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

func ifOriginalURLAlreadySavedLocally(originalUrl string) string {
	for key, val := range allUrls {
		if val == originalUrl {
			return key
		}
	}
	return ""
}

func ifShortURLAlreadySavedLocally(shortURL string) bool {
	_, exists := allUrls[shortURL]
	return exists
}

func saveOriginalURLInDB(originalUrl string) (string, error) {
	shortURL, err := ifOriginalURLAlreadySavedInDB(originalUrl)
	if err != nil {
		return "", err
	}

	// If originalUrl already saved locally, return existing shortURL
	if shortURL != "" {
		return shortURL, nil
	}

	// If originalUrl is not saved, generate new shortURL
	shortURL = utils.GenerateShortURL(originalUrl)

	err = saveShortURLInDB(shortURL, originalUrl)
	if err != nil {
		return "", errors.New("Error: the generated short URL already exists")
	}

	return shortURL, nil
}

func ifOriginalURLAlreadySavedInDB(originalUrl string) (string, error) {
	var shortURL string

	data, err := db.Query(`
			SELECT short FROM urls
			WHERE original = $1;
		`, originalUrl)

	for data.Next() {
		err = data.Scan(&shortURL)
		if err != nil {
			return "", errors.New("Error: scanning results after getting the original URL form the DB by short URL")
		}
	}

	return shortURL, nil
}

func saveShortURLInDB(shortURL string, originalUrl string) error {
	_, err := db.Query(`
			INSERT INTO urls (short, original)
			VALUES ($1, $2);
		`, shortURL, originalUrl)

	return err
}
