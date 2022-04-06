package db

import (
	"errors"
	"fmt"

	"github.com/Arcady1/OZON_Internship_API/utils"
)

var allUrls = make(map[string]string)

func SaveURLInDB(originalUrl string) (string, error) {
	shortUrl := utils.GenerateShortURL(originalUrl)

	_, exists := allUrls[shortUrl]
	if exists == false {
		allUrls[shortUrl] = originalUrl
		fmt.Println("allUrls:", allUrls)
	} else {
		return "", errors.New("The short URL already exists")
	}

	return shortUrl, nil
}
