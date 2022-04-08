package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Arcady1/OZON_Internship_API/controllers"
	_ "github.com/Arcady1/OZON_Internship_API/models"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1.0/url", controllers.GetOriginalURL).Methods("GET")
	router.HandleFunc("/api/v1.0/url", controllers.PostOriginalURL).Methods("POST")

	PORT := os.Getenv("PORT")

	log.Printf("Server started at: %v\n", PORT)

	err := http.ListenAndServe(":"+PORT, router)
	if err != nil {
		log.Println(err)
	}
}
