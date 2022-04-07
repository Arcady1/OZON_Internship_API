package main_test

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/Arcady1/OZON_Internship_API/models"
	"github.com/Arcady1/OZON_Internship_API/utils"
)

var db *sql.DB

type response = utils.JsonResponse

var hostURL string = "http://127.0.0.1:8000/api/v1.0"

func TestMain(m *testing.M) {
	var err error
	db, err = models.GetDB()

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("\n\t  Success: connected to the DB")
	}

	utils.SetDataStorageIsDB(false)
	m.Run()

	utils.SetDataStorageIsDB(true)
	m.Run()
}

func TestGetNonExistentShortURL(t *testing.T) {
	resp := &response{}
	getJson(t, hostURL+"/url?short=lasd9p21_X", resp)

	checkResponseCode(t, resp.Status, 500)

	if resp.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", resp.Data)
	}
}

func TestGetShortURLWithManyParams(t *testing.T) {
	resp := &response{}
	getJson(t, hostURL+"/url?short=lasd9p21_X&add=123", resp)

	checkResponseCode(t, resp.Status, 400)

	if resp.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", resp.Data)
	}
}

func TestGetShortURLNoShortURLParam(t *testing.T) {
	resp := &response{}
	getJson(t, hostURL+"/url?add=123", resp)

	checkResponseCode(t, resp.Status, 400)

	if resp.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", resp.Data)
	}
}

func TestGetShortURLShortURLIsEmpty(t *testing.T) {
	resp := &response{}
	getJson(t, hostURL+"/url?short=", resp)

	checkResponseCode(t, resp.Status, 400)

	if resp.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", resp.Data)
	}
}

func TestPostOriginalURLWithManyParams(t *testing.T) {
	resp := &response{}

	postJson(t, hostURL+"/url?original=https://google.com&add=test", resp)

	checkResponseCode(t, resp.Status, 400)

	if resp.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", resp.Data)
	}
}

func TestPostOriginalURLNoOriginalURLParam(t *testing.T) {
	resp := &response{}
	postJson(t, hostURL+"/url?add=test", resp)

	checkResponseCode(t, resp.Status, 400)

	if resp.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", resp.Data)
	}
}

func TestPostOriginalURLOriginalURLIsEmpty(t *testing.T) {
	resp := &response{}
	postJson(t, hostURL+"/url?original=", resp)

	checkResponseCode(t, resp.Status, 400)

	if resp.Data != nil {
		t.Errorf("Expected an empty data. Got %v\n", resp.Data)
	}
}

func getJson(t *testing.T, url string, target interface{}) {
	resp, err := http.Get(url)

	if err != nil {
		t.Errorf("Impossible to make GET request. URL: %s\n", url)
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(target)
}

func postJson(t *testing.T, url string, target interface{}) {
	resp, err := http.PostForm(url, nil)

	if err != nil {
		t.Errorf("Impossible to make POST request. URL: %s\n", url)
	}

	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		t.Errorf("Impossible to read the POST response body. URL: %s\n", url)
	}

	errUnmarshal := json.Unmarshal(body, target)
	if errUnmarshal != nil {
		t.Errorf("Impossible to transform the POST response body into response structure. URL: %s\n", url)
	}
}

func checkResponseCode(t *testing.T, realStatusCode uint16, expectedStatusCode uint16) {
	if realStatusCode != expectedStatusCode {
		t.Errorf("Expected response code: %v. Got %v\n", expectedStatusCode, realStatusCode)
	}
}
