package controllers

import (
	"fmt"
	"net/http"

	"github.com/Arcady1/OZON_Internship_API/db"
	"github.com/Arcady1/OZON_Internship_API/utils"
)

func GetShortURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetShortURL")

	params := r.FormValue("short")
	fmt.Println("params", params)
	message := "Getting a short URL"
	response := &utils.JsonResponse{Status: 200, Message: message}

	utils.ResponseWriter(w, response)
}

func PostOriginalURL(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostOriginalURL")

	response := &utils.JsonResponse{}
	paramName := "original"

	originalUrl, err := utils.CheckErr(w, r, paramName, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	shortUrl, err := db.SaveURLInDB(originalUrl)
	if err != nil {
		fmt.Println(err)
		response.Status = 500
		response.Message = "The short URL already exists"
		utils.ResponseWriter(w, response)

		return
	} else {
		response.Status = 201
		response.Message = "Original URL is saved"
		response.Data = map[string]string{"shortURL": shortUrl}
		utils.ResponseWriter(w, response)
	}
}
