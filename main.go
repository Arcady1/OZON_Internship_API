package main

import (
	"log"
	"os"

	_ "github.com/Arcady1/OZON_Internship_API/models"
)

func main() {
	a := App{}
	a.Initialize(false)

	PORT := os.Getenv("PORT")
	log.Printf("Server started at: %v\n", PORT)

	a.Run(PORT)
}
