package main

import (
	"Phonebook/data"
	"Phonebook/handlers"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"Phonebook/db"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], ":\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "\n")
}

func main() {
	// основные настройки к базе
	flag.Usage = Usage
	server := flag.String("s", "localhost", "serverDB")
	port := flag.String("p", "5432", "port")
	PostgresUser := flag.String("u", "postgres", "Database user")
	PostgresPassword := flag.String("pass", "admin", "DB password")
	PostgresDB := flag.String("d","postgres","Name Database")

	flag.Parse()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", *PostgresUser, *PostgresPassword, *server, *port,*PostgresDB)
	fmt.Println(dsn)
	postgresrepo, err:=db.NewPostgresrepo(&dsn)
	if err !=nil  {
		fmt.Println("Error DB. Please check your connect for DB",err)
		log.Fatal()
	}
	db.SetRepository(postgresrepo)
	err = postgresrepo.Db.Ping()
	if err !=nil  {
		fmt.Println("Error DB. Please check your connect for DB",err)
		log.Fatal()
	}
	postgresrepo.Create()


	datarepo := data.NewDataRepo()
	err = datarepo.GetCountryName()
	if err !=nil  {
		fmt.Println("Source unreachable",err)
	}
	err = datarepo.GetPhoneCode()
	if err !=nil  {
		fmt.Println("Source unreachable",err)
	}
	err = db.Insert(*datarepo)

	if err !=nil  {
		fmt.Println("Error DB. Please check your connect for DB",err)
	}

	handlersT := handlers.NewHandlerT()
	rT := mux.NewRouter()
	rT.HandleFunc("/reload", handlersT.Reload).Methods("POST")
	rT.HandleFunc("/code/{[]}", handlersT.SelectCountry).Methods("GET")
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", rT)
}
