package models

import (
	"fmt"

	"github.com/Arcady1/OZON_Internship_API/utils"
)

var allUrls = make(map[string]string)

func SaveURLInDB(originalUrl string) (string, error) {
	db, err := GetDB()

	if err != nil {
		return "", err
	}

	shortUrl := utils.GenerateShortURL(originalUrl)

	// TODO Save locally
	// _, exists := allUrls[shortUrl]
	// if exists == false {
	// 	allUrls[shortUrl] = originalUrl
	// 	fmt.Println("allUrls:", allUrls)
	// } else {
	// 	return "", errors.New("The short URL already exists")
	// }

	// ? Save into the Database
	rows, err := db.Query("SELECT * FROM urls")
	if err != nil {
		return "", err
	}
	fmt.Println("rows:\n", rows)

	return shortUrl, nil
}
