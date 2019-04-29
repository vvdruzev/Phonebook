package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
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
	PostgresUser := flag.String("u", "", "Database user")
	PostgresPassword := flag.String("pass", "", "DB password")
	PostgresDB := flag.String("d","","Name Database")

	flag.Parse()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", *PostgresUser, *PostgresPassword, *server, *port,*PostgresDB)
	postgresrepo, err:=NewPostgresrepo(&dsn)
	if err !=nil  {
		fmt.Println("Error DB. Please check your connect for DB",err)
		log.Fatal()
	}
	err = postgresrepo.db.Ping()
	if err !=nil  {
		fmt.Println("Error DB. Please check your connect for DB",err)
		log.Fatal()
	}

	postgresrepo.Create()
	countryname,err := GetData(COUNTRYNAME)
	if err !=nil  {
		fmt.Println("Source unreachable",err)
	}
	phonecode, err := GetData(PHONECODE)
	if err !=nil  {
		fmt.Println("Source unreachable",err)
	}
	err = postgresrepo.Insert(countryname,phonecode)

	if err !=nil  {
		fmt.Println("Error DB. Please check your connect for DB",err)
	}

	handlersT := NewHandlerT()
	handlersT.GetConn(postgresrepo)

	rT := mux.NewRouter()
	rT.HandleFunc("/reload", handlersT.Reload).Methods("POST")
	rT.HandleFunc("/code/{[]}", handlersT.SelectCountry).Methods("GET")
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8081", rT)
}
