package models

import (
	"database/sql"
	"log"
)

func GenerateSampleData(db *sql.DB) {
	generateLocalData()
	generateDBData(db)
}

func generateLocalData() {
	allUrls["lasd9p21_X"] = "https://google.com"
	allUrls["lk9_aslZxa"] = "https://yandex.ru"
}

func generateDBData(db *sql.DB) {
	var err error

	for key, val := range allUrls {
		_, err = db.Exec(`
			INSERT INTO urls (short, original)
			VALUES ($1, $2);
		`, key, val)

		if err != nil {
			log.Fatalf("Impossible to generate sample data: %v\n", err)
		}
	}
}
