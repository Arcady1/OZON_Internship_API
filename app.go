package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Arcady1/OZON_Internship_API/models"
	"github.com/Arcady1/OZON_Internship_API/utils"
	_ "github.com/Arcady1/OZON_Internship_API/utils"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(testMode bool) {
	var err error

	a.DB, err = models.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initRoutes()

	if testMode == true {
		a.generateSampleData()
	}
}

func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(":"+port, a.Router))

	err := models.Migrate()
	if err != nil {
		log.Fatalf("Error on Migrate(). Err: %v\n", err)
	}
}

func (a *App) getOriginalURL(w http.ResponseWriter, r *http.Request) {
	response := &utils.JsonResponse{}
	paramName := "short"

	shortURL, err := utils.CheckQuery(w, r, paramName, 1)
	if err != nil {
		log.Println(err)
		return
	}

	originalUrl, err := models.GetOriginalURL(shortURL)
	if err != nil {
		log.Println(err)
		response.Status = 500
		response.Message = fmt.Sprintf("%v", err)
		utils.ResponseWriter(w, response, response.Status)

		return
	} else {
		response.Status = 200
		response.Message = "Getting the original URL"
		response.Data = map[string]string{"originalUrl": originalUrl}
		utils.ResponseWriter(w, response, response.Status)
	}
}

func (a *App) postOriginalURL(w http.ResponseWriter, r *http.Request) {
	response := &utils.JsonResponse{}
	paramName := "original"

	originalUrl, err := utils.CheckQuery(w, r, paramName, 1)
	if err != nil {
		log.Println(err)
		return
	}

	shortUrl, err := models.SaveOriginalURL(originalUrl)
	if err != nil {
		log.Println(err)
		response.Status = 500
		response.Message = fmt.Sprintf("%v", err)
		utils.ResponseWriter(w, response, response.Status)

		return
	} else {
		response.Status = 201
		response.Message = "Original URL is saved"
		response.Data = map[string]string{"shortURL": shortUrl}
		utils.ResponseWriter(w, response, response.Status)
	}
}

func (a *App) generateSampleData() {
	models.GenerateSampleData(a.DB)
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/api/v1.0/url", a.getOriginalURL).Methods("GET")
	a.Router.HandleFunc("/api/v1.0/url", a.postOriginalURL).Methods("POST")
}
