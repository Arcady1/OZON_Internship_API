package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseWriter(w http.ResponseWriter, data JsonResponse) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
