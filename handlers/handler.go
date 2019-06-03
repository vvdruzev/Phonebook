package handlers

import (
	"Phonebook/data"
	"net/http"
	"Phonebook/db"
	"Phonebook/util"
	"github.com/gorilla/mux"
	"Phonebook/logger"
	"github.com/sirupsen/logrus"
)

type Handler struct {
}

func NewHandler() *Handler  {
	return &Handler{
	}
}

func (h Handler) Reload (w http.ResponseWriter, r *http.Request) {
	datarepo := data.NewDataRepo()
	logger.WithFields(logrus.Fields{"method":r.Method,}).Info("Reload data")
	err := datarepo.GetCountryName()
	if err !=nil  {
		logger.Error("Source unreachable ",err)
	}
	err = datarepo.GetPhoneCode()
	if err !=nil  {
		logger.Error("Source unreachable ",err)
	}
	logger.Info("insert data in DB")
	err=db.Reload(*datarepo)
	if err != nil {
		logger.Error("Error connect to DB",err)
	}
}

func (h Handler) SelectCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country := vars["country"]
	logger.WithFields(logrus.Fields{"method":r.Method,}).Info("Select Country ",country)

	rows, err :=db.Select(country)

	if err != nil {
		logger.Error(err)
		util.ResponseError(w,http.StatusNotFound,"404 page not found")
		return
	}
	util.ResponseOk(w,rows)
	logger.Debug("Found Country ",country)
}
