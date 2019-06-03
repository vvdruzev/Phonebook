package main

import (
	"Phonebook/data"
	"Phonebook/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"Phonebook/schema"
	"Phonebook/db"
	"Phonebook/logger"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"encoding/json"
)

type PostgresrepoTest struct {

}

func (d PostgresrepoTest) Select (country string ) (schema.ResponseCode,error) {
	var p schema.ResponseCode
	if strings.ToUpper(country)== "JAMAICA" {
		p = schema.ResponseCode{
			"+1-876",
		}
	}else{
		return p, errors.New("not found")
	}
	return p,nil
}

func (d PostgresrepoTest) Reload (repo data.DataRepo) error {
	return nil
}

func (d PostgresrepoTest) Insert(repo data.DataRepo) error {
	return nil
}

func (d PostgresrepoTest) Create() error {
	return nil
}

func (d PostgresrepoTest) Close()  {}

func NewPostgresrepoTest() *PostgresrepoTest {
	return &PostgresrepoTest{}
}

type TestCase struct {
	countryName string
	Result  SearchResponse
	errTest string
}

type SearchResponse struct {
	status int
	response string
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("handler returned wrong status code: got %d  want %d\n", actual, expected)
	}
}

func TestNotFoundCountry(t *testing.T) {
	logger.NewLogger()
	item := []TestCase{
		{
			countryName: "Jama",
			Result: SearchResponse{
				status:   404,
				response: "404 page not found",
			},

		},
	}

	postgresrepoTest := NewPostgresrepoTest()
	h := handlers.NewHandler()
	db.SetRepository(postgresrepoTest)

	for _,val := range item {

		req, err := http.NewRequest("GET", "/code/", nil)
		if err != nil {
			t.Fatal(err)
		}
		vars := map[string]string{
			"country": val.countryName,
		}
		req = mux.SetURLVars(req,vars)
		handler := http.HandlerFunc(h.SelectCountry)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		checkResponseCode(t, http.StatusNotFound, rr.Code)

		expected := val.Result.response
		var m map[string]string
		json.Unmarshal(rr.Body.Bytes(), &m)

		if m["error"] != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				m["error"], expected)
		}
	}
}


func TestFoundCountry(t *testing.T) {
	logger.NewLogger()
	item := []TestCase{
		{
			countryName: "jamaica",
			Result: SearchResponse{
				status:   http.StatusOK,
				response: "+1-876",
			},
		},{
			countryName:"jamaiCa",
			Result: SearchResponse{
				status:   http.StatusOK,
				response: "+1-876",
			},

		},
	}

	postgresrepoTest := NewPostgresrepoTest()
	h := handlers.NewHandler()
	db.SetRepository(postgresrepoTest)

	for _,val := range item {

		req, err := http.NewRequest("GET", "/code/", nil)
		if err != nil {
			t.Fatal(err)
		}

		vars := map[string]string{
			"country": val.countryName,
		}
		req = mux.SetURLVars(req,vars)
		handlerT := http.HandlerFunc(h.SelectCountry)
		rr := httptest.NewRecorder()
		handlerT.ServeHTTP(rr, req)

		checkResponseCode(t, http.StatusOK, rr.Code)

		expected := val.Result.response

		var m map[string]string
		json.Unmarshal(rr.Body.Bytes(), &m)

		if m["PhoneCode"] != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				m["Phonebook"], expected)
		}
	}
}