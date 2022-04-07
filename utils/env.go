package utils

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	envPath, err := filepath.Abs(".env")

	if err != nil {
		log.Fatalln("Error: getting the .env file path")
	}

	err = godotenv.Load(envPath)

	if err != nil {
		log.Fatalln("Impossible to read .env file")
	}
}
