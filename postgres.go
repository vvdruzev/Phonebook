package main

import (
	"database/sql"
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

)

type Postgresrepo struct {
	db *sql.DB
}

type productT struct {

	CountryName string
	PhoneCode string
	CountryCode string
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

func (d *Postgresrepo) Close() {
	d.db.Close()
}

func (d Postgresrepo) Create() error {
	_, err := d.db.Exec("DROP TABLE IF EXISTS PhoneBook;")
	if err != nil {
		return err
	}
	_, err = d.db.Exec(`
	CREATE TABLE PhoneBook (
		CountryCode varchar(10) NOT NULL,
		CountryName varchar(45) DEFAULT NULL,
		PhoneCode varchar(45) DEFAULT NULL,
		PRIMARY KEY (CountryCode));`)
	if err != nil {
		return err

	}
	return nil
}

func (d Postgresrepo) Insert(country map[string]interface{},phone map[string]interface{}) error {
	sqlStr := "INSERT INTO PhoneBook(CountryCode, CountryName, PhoneCode) VALUES ($1,$2,$3)"

	tx, err:=d.db.Begin()
	for key, val := range country {
		_,err = tx.Exec(sqlStr,key, fmt.Sprintf("%s",val), fmt.Sprintf("%s",phone[key]))
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}

func (d Postgresrepo) Reload (country map[string]interface{},phone map[string]interface{}) error  {

	_, err := d.db.Exec("truncate PhoneBook")
	if err != nil {
		return err
	}
	err = d.Insert(country,phone)
	if err != nil {
		return err
	}
	return nil
}

func (d Postgresrepo) Select (country string ) ([]productT,error) {
	rows, err := d.db.Query("select * from PhoneBook where upper(CountryName)=$1 LIMIT 1", strings.ToUpper(country))
	defer rows.Close()
	if err != nil {
		return nil,err
	}
	var products []productT
	for rows.Next() {
		p := productT{}
		err := rows.Scan(&p.CountryCode, &p.CountryName, &p.PhoneCode)
		if err != nil {
			return nil,err
		}
		products = append(products, p)
	}

	return products,nil
}