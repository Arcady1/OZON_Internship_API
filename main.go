package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	// "log"
)

type JsonResponse struct {
	Status  uint16 `json:"status"`
	Message string `json:"message"`
}

type Urls struct {
	Short    string `json:"shor"`
	Original string `json:"original"`
}

var allUrls = []Urls{}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1.0/url", GetShortURL).Methods("GET")
	// router.HandleFunc("/api/v1.0/url", PostOriginalURL).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	fmt.Printf("Server started at %v\n", PORT)

	http.ListenAndServe(":"+PORT, router)
}

func GetShortURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetShortURL")

	params := mux.Vars(r)
	fmt.Println("params", params)
	message := "Nice!"
	response := JsonResponse{Status: 200, Message: message}

	utils.ResponseWriter(w, response)
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(response)
}
