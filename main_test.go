package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Arcady1/OZON_Internship_API/utils"
)

var db *sql.DB

var response = utils.JsonResponse{}

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize(true)

	log.Println("TESTS WHEN DATA STORAGE IS LOCALSTORAGE")
	utils.SetDataStorageIsDB(false)
	m.Run()

	log.Println("TESTS WHEN DATA STORAGE IS DATABASE")
	utils.SetDataStorageIsDB(true)
	m.Run()
}

func TestGetNonExistentShortURL(t *testing.T) {
	var url = "/url?short=abcd1e23_F"

	fillResponse(t, url, "GET")
	checkResponseCode(t, response.Status, 500)

	if response.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", response.Data)
	}
}

func TestGetShortURLWithManyParams(t *testing.T) {
	var url = "/url?short=lasd9p21_X&add=123"

	fillResponse(t, url, "GET")
	checkResponseCode(t, response.Status, 400)

	if response.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", response.Data)
	}
}

func TestGetShortURLNoShortURLParam(t *testing.T) {
	var url = "/url?add=123"

	fillResponse(t, url, "GET")
	checkResponseCode(t, response.Status, 400)

	if response.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", response.Data)
	}
}

func TestGetShortURLShortURLIsEmpty(t *testing.T) {
	var url = "/url?short="

	fillResponse(t, url, "GET")
	checkResponseCode(t, response.Status, 400)

	if response.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", response.Data)
	}
}

func TestGetShortURCorrect(t *testing.T) {
	var url = "/url?short=lk9_aslZxa"

	fillResponse(t, url, "GET")
	checkResponseCode(t, response.Status, 200)

	if response.Data == nil {
		t.Errorf("Expected non empty data. Got %v\n", response.Data)
	}
}

func TestPostOriginalURLWithManyParams(t *testing.T) {
	var url = "/url?original=https://google.com&add=test"

	fillResponse(t, url, "POST")
	checkResponseCode(t, response.Status, 400)

	if response.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", response.Data)
	}
}

func TestPostOriginalURLNoOriginalURLParam(t *testing.T) {
	var url = "/url?add=test"

	fillResponse(t, url, "POST")
	checkResponseCode(t, response.Status, 400)

	if response.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", response.Data)
	}
}

func TestPostOriginalURLOriginalURLIsEmpty(t *testing.T) {
	var url = "/url?original="

	fillResponse(t, url, "POST")
	checkResponseCode(t, response.Status, 400)

	if response.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", response.Data)
	}
}

func TestPostOriginalCorrect(t *testing.T) {
	var url = "/url?original=https://google.com"

	fillResponse(t, url, "POST")
	checkResponseCode(t, response.Status, 201)

	if response.Data == nil {
		t.Errorf("Expected non empty data. Got %v\n", response.Data)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func fillResponse(t *testing.T, url string, method string) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Errorf("Impossible to make %v request. Err: %v\n", method, err)
	}

	serverResponse := executeRequest(req)

	err = json.Unmarshal(serverResponse.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Impossible to parse JSON into response. Err: %v\n", err)
	}
}

func checkResponseCode(t *testing.T, realStatusCode uint16, expectedStatusCode uint16) {
	if realStatusCode != expectedStatusCode {
		t.Errorf("Expected response code: %v. Got %v\n", expectedStatusCode, realStatusCode)
	}
}
