package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/Arcady1/OZON_Internship_API/controllers"
	_ "github.com/Arcady1/OZON_Internship_API/models"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1.0/url", controllers.GetOriginalURL).Methods("GET")
	router.HandleFunc("/api/v1.0/url", controllers.PostOriginalURL).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	fmt.Printf("Server started at %v\n", PORT)

	err := http.ListenAndServe(":"+PORT, router)
	if err != nil {
		fmt.Println(err)
	}
}
