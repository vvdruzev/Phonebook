package handlers

import (
	"Phonebook/data"
	"fmt"
	"io"
	"net/http"
	"strings"
	"Phonebook/db"
)

type HandlerT struct {
}

func NewHandlerT() *HandlerT  {
	return &HandlerT{
	}
}

func (h HandlerT) Reload (w http.ResponseWriter, r *http.Request) {
	datarepo := data.NewDataRepo()
	err := datarepo.GetCountryName()
	if err !=nil  {
		fmt.Println("Source unreachable",err)
	}
	err = datarepo.GetPhoneCode()
	if err !=nil  {
		fmt.Println("Source unreachable",err)
	}
	err=db.Reload(*datarepo)
	if err != nil {
		fmt.Println("DB Error",err)
	}
}

func (h HandlerT) SelectCountry(w http.ResponseWriter, r *http.Request) {
	country:=strings.TrimLeft(r.URL.String(),"/code/")

	rows, err :=db.Select(country)
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
