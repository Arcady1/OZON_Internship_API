package models

func ifOriginalURLAlreadySavedLocally(originalUrl string) string {
	for key, val := range allUrls {
		if val == originalUrl {
			return key
		}
	}
	return ""
}

func ifShortURLAlreadySavedLocally(shortURL string) (exists bool) {
	_, exists = allUrls[shortURL]
}
