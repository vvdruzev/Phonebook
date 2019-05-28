package main

import (
	"Phonebook/data"
	"Phonebook/handlers"
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"Phonebook/schema"
	"Phonebook/db"
	"Phonebook/logger"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type PostgresrepoTest struct {
	DB *sql.DB
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
func (d PostgresrepoTest) Close()  {

}

func NewPostgresrepoTest() *PostgresrepoTest {
	return &PostgresrepoTest{
	}
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

func TestSelectCountry(t *testing.T) {
	logger.NewLogger()
	item := []TestCase{
		{
		countryName: "jamaica",
		Result: SearchResponse{
			status:   200,
			response: "{\"PhoneCode\":\"+1-876\"}\n",
		},
	},{
		countryName:"jamaiCa",
		Result: SearchResponse{
			status:   200,
			response: "{\"PhoneCode\":\"+1-876\"}\n",
		},

	},
	{
		countryName: "Jama",
		Result: SearchResponse{
			status:   404,
			response: "{\"error\":\"404 page not found\"}\n",
		},

	},

	}
	postgresrepoTest := NewPostgresrepoTest()
	h := handlers.NewHandlerT()
	db.SetRepository(postgresrepoTest)

	for _,val := range item {

		rr := httptest.NewRecorder()
		fmt.Println(val)

		req, err := http.NewRequest("GET", "/code/", nil)
		if err != nil {
			t.Fatal(err)
		}
		vars := map[string]string{
			"country": val.countryName,
		}
		req = mux.SetURLVars(req,vars)
		handler := http.HandlerFunc(h.SelectCountry)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != val.Result.status {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, val.Result.status)
		}

		expected := val.Result.response

		if fmt.Sprintf(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}