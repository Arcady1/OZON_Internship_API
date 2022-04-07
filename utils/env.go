package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	envPath, err := filepath.Abs(".env")

	if err != nil {
		fmt.Println("Error: getting the .env file path")
		os.Exit(2)
	}

	err = godotenv.Load(envPath)

	if err != nil {
		fmt.Println("\nImpossible to read .env file")
		os.Exit(1)
	}
}
