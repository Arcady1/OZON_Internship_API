package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/Arcady1/OZON_Internship_API/controllers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1.0/url", controllers.GetShortURL).Methods("GET")
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
