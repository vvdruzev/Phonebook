package main

import (
	"Phonebook/data"
	"Phonebook/handlers"
	"Phonebook/db"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"

	"fmt"
	"log"
	"net/http"
	"time"
	"Phonebook/logger"
)


type Config struct {
	PostgresDB           string `envconfig:"POSTGRES_DB"`
	PostgresUser         string `envconfig:"POSTGRES_USER"`
	PostgresPassword     string `envconfig:"POSTGRES_PASSWORD"`
	ProxyUrl			 string `envconfig:"HTTP_PROXY"`
}



func main() {
	logger.NewLogger()
	// основные настройки к базе
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		logger.Error(err)
		log.Fatal(err)
	}


	dsn := fmt.Sprintf("postgres://%s:%s@postgres/%s?sslmode=disable",cfg.PostgresUser, cfg.PostgresPassword , cfg.PostgresDB)
	postgresrepo, err:=db.NewPostgresrepo(&dsn)
	if err !=nil  {
		logger.Error("Error DB. Please check your connect for DB",err,dsn)
		log.Fatal()
	}
	for {
		db.SetRepository(postgresrepo)
		err = postgresrepo.Db.Ping()
		if err != nil {
			logger.Error("Error DB. Please check your connect for DB",err,dsn)
			time.Sleep(1000)
		}else {
			break
		}
	}
	if err != nil {
		logger.Error("Error DB. Please check your connect for DB",err,dsn)
		log.Fatal()
	}

	defer db.Close()

	logger.Info("Connect to DB ",cfg.PostgresDB, ", user " , cfg.PostgresUser)

	data.SetClient(cfg.ProxyUrl)
	logger.Debug("Set proxy:", cfg.ProxyUrl)

	handlers := handlers.NewHandler()
	r := mux.NewRouter()
	r.HandleFunc("/reload", handlers.Reload).Methods("POST")
	r.HandleFunc("/code/{country}", handlers.SelectCountry).Methods("GET")
	logger.Info("starting server at :8080")
	http.ListenAndServe(":8080", r)
}
