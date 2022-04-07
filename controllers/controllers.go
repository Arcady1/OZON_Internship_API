package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Arcady1/OZON_Internship_API/models"
	"github.com/Arcady1/OZON_Internship_API/utils"
)

func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	response := &utils.JsonResponse{}
	paramName := "short"

	shortURL, err := utils.CheckQuery(w, r, paramName, 1)
	if err != nil {
		log.Println(err)
		return
	}

	originalUrl, err := models.GetURL(shortURL)
	if err != nil {
		log.Println(err)
		response.Status = 500
		response.Message = fmt.Sprintf("%v", err)
		utils.ResponseWriter(w, response)

		return
	} else {
		response.Status = 200
		response.Message = "Getting the original URL"
		response.Data = map[string]string{"originalUrl": originalUrl}
		utils.ResponseWriter(w, response)
	}
}

func PostOriginalURL(w http.ResponseWriter, r *http.Request) {
	response := &utils.JsonResponse{}
	paramName := "original"

	originalUrl, err := utils.CheckQuery(w, r, paramName, 1)
	if err != nil {
		log.Println(err)
		return
	}

	shortUrl, err := models.SaveURL(originalUrl)
	if err != nil {
		log.Println(err)
		response.Status = 500
		response.Message = fmt.Sprintf("%v", err)
		utils.ResponseWriter(w, response)

		return
	} else {
		response.Status = 201
		response.Message = "Original URL is saved"
		response.Data = map[string]string{"shortURL": shortUrl}
		utils.ResponseWriter(w, response)
	}
}
