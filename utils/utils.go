package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Status  uint16      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseWriter(w http.ResponseWriter, data *JsonResponse) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*data)
}

func CheckQuery(w http.ResponseWriter, r *http.Request, paramName string, paramNumber int) (string, error) {
	response := &JsonResponse{}
	paramVal := r.FormValue(paramName)
	queryParamsMap := r.URL.Query()

	if len(queryParamsMap) != paramNumber {
		response.Status = 400
		response.Message = "Wrong number of query parameters"
		ResponseWriter(w, response)
		return "", errors.New(response.Message)
	}

	_, exists := queryParamsMap[paramName]
	if exists == false {
		response.Status = 400
		response.Message = fmt.Sprintf("Parameter '%v' is missing", paramName)
		ResponseWriter(w, response)
		return "", errors.New(response.Message)
	}

	if paramVal == "" {
		response.Status = 400
		response.Message = fmt.Sprintf("Parameter '%v' is empty", paramName)
		ResponseWriter(w, response)
		return "", errors.New(response.Message)
	}

	return paramVal, nil
}

func GenerateShortURL(originalURL string) string {
	// num := int(time.Now().UnixNano() / int64(time.Millisecond))
	// shortURL := originalURL + strconv.Itoa(num)
	shortURL := "lasd9p21_X"

	fmt.Println("shortURL:", shortURL)
	return shortURL
}
