package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type PostgresrepoTest struct {
	DB *sql.DB
}

func (d PostgresrepoTest) Select (country string ) ([]productT,error) {
	var products []productT
	var p productT
	if strings.ToUpper(country)== "JAMAICA" {
		p = productT{
			CountryName: "Jamaica",
			PhoneCode:   "+1-876",
		}
	}else{
		return nil,nil
	}
	products =  append(products,p)
	return products,nil
}

func (d PostgresrepoTest) Reload (map[string]interface{},map[string]interface{}) error {
	return nil
}

func (d PostgresrepoTest) Insert(map[string]interface{},map[string]interface{}) error {
	return nil
}

func (d PostgresrepoTest) Create() error {
	return nil
}

func NewPostgresrepoTest() *PostgresrepoTest {
	return &PostgresrepoTest{
	}
}

type TestCase struct {
	req string
	Result  SearchResponse
	errTest string
}

type SearchResponse struct {
	status int
	response string
}

func TestSelectCountry(t *testing.T) {
	item := []TestCase{
		{
		req: "/code/jamaica",
		Result: SearchResponse{
			status:   200,
			response: `{"Jamaica": "+1-876"}`,
		},
	},{
		req:"/code/jamaiCa",
		Result: SearchResponse{
			status:   200,
			response: `{"Jamaica": "+1-876"}`,
		},

	},
	{
		req: "/code/Jama",
		Result: SearchResponse{
			status:   404,
			response: `404 page not found`,
		},

	},

	}
	postgresrepoTest := NewPostgresrepoTest()
	h := NewHandlerT()
	h.Postgresrepo=postgresrepoTest

	for _,val := range item {

		rr := httptest.NewRecorder()
		fmt.Println(val)

		req, err := http.NewRequest("GET", val.req, nil)
		if err != nil {
			t.Fatal(err)
		}
		handler := http.HandlerFunc(h.SelectCountry)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != val.Result.status {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, val.Result.status)
		}

		expected := val.Result.response
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}