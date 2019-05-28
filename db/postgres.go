package db

import (
	"Phonebook/data"
	"database/sql"
	"fmt"
	"strings"
	_ "github.com/lib/pq"
	"Phonebook/schema"
)

type Postgresrepo struct {
	Db *sql.DB
}


func NewPostgresrepo(dsn *string) (*Postgresrepo,error) {
	db, err := sql.Open("postgres", *dsn)
	if err != nil {
		return nil, err
	}
	return &Postgresrepo{
		db,
	}, nil
}


type dbError struct {
	method string
	Err error
}


func (re *dbError) Error() string {
	return fmt.Sprintf(
		"DB error:  %s, err: %v",
		re.method,
		re.Err,
	)
}

func (db Postgresrepo) Close() {
	db.Db.Close()
}

func (db Postgresrepo) Insert(data data.DataRepo) error {
	sqlStr := "INSERT INTO PhoneBook(CountryCode, CountryName, PhoneCode) VALUES ($1,$2,$3)"
	country := data.CountryName
	phone := data.PhoneCode
	tx, err:=db.Db.Begin()
	for key, val := range country {
		_,err = tx.Exec(sqlStr,key, fmt.Sprintf("%s",val), fmt.Sprintf("%s",phone[key]))
		if err != nil {
			return &dbError{method: sqlStr, Err: err}
		}
	}
	tx.Commit()
	return nil
}

func (d Postgresrepo) Reload (data data.DataRepo) error  {

	_, err := d.Db.Exec("truncate PhoneBook")
	if err != nil {
		return err
	}
	err = d.Insert(data)
	if err != nil {
		return err
	}
	return nil
}

func (d Postgresrepo) Select (country string ) (schema.ResponseCode,error) {
	sqlStr := "select PhoneCode from PhoneBook where upper(CountryName)=$1 LIMIT 1"
	rows:= d.Db.QueryRow(sqlStr, strings.ToUpper(country))
	p := schema.ResponseCode{}
	err := rows.Scan(&p.PhoneCode)
	if err == sql.ErrNoRows {
		return p, &dbError{method: "Not found", Err: err}
	} else if err != nil {
		return 	p, &dbError{method: sqlStr, Err: err}
	}

	return p,nil
}