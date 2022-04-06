package db

import "fmt"

// type urls struct {
// 	Short    string `json:"short"`
// 	Original string `json:"original"`
// }

var allUrls = make(map[string]string)

func SaveURLInDB(originalUrl string) (string, uint8) {
	shortUrl := "http://short_url.com"
	_, exists := allUrls[shortUrl]

	if exists == false {
		allUrls[shortUrl] = originalUrl
		fmt.Println("allUrls:", allUrls)
	} else {
		return "", 1
	}

	return shortUrl, 0
}
