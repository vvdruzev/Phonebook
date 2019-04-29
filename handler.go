package main

import (
	"io"
	"net/http"
	"strings"
)

type HandlerT struct {
	Postgresrepo Conn
}

type Conn interface {
	Create() error
	Reload (map[string]interface{},map[string]interface{}) error
	Select(country string ) ([]productT,error)
	Insert(map[string]interface{},map[string]interface{}) error
}

func NewHandlerT() *HandlerT  {
	return &HandlerT{
	}
}

func (h *HandlerT) GetConn (d *Postgresrepo) *HandlerT {
	h.Postgresrepo = *d
	return h
}

func (h HandlerT) Reload (w http.ResponseWriter, r *http.Request) {
	countryname,err := GetData(COUNTRYNAME)
	if err !=nil  {
		panic(err)
	}
	phonecode, err := GetData(PHONECODE)
	if err !=nil  {
		panic(err)
	}
	err=h.Postgresrepo.Reload(countryname,phonecode)
	if err != nil {

	}
}

func (h HandlerT) SelectCountry(w http.ResponseWriter, r *http.Request) {
	country:=strings.TrimLeft(r.URL.String(),"/code/")

	rows, err :=h.Postgresrepo.Select(country)
	if err != nil {

	}
	if rows==nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "404 page not found")
		return
	}
	w.WriteHeader(http.StatusOK)
	res := "{"
	for _,val:= range rows{
		res += "\"" + val.CountryName +"\": \""+ val.PhoneCode +"\","
	}
	res = strings.TrimRight(res,",")+ "}"
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, res)
}
