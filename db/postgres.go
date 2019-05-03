package db

import (
	"Phonebook/data"
	"database/sql"
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
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

func (d *Postgresrepo) Close() {
	d.Db.Close()
}

func (d Postgresrepo) Create() error {
	_, err := d.Db.Exec("DROP TABLE IF EXISTS PhoneBook;")
	if err != nil {
		return err
	}
	_, err = d.Db.Exec(`
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

func (db Postgresrepo) Insert(data data.DataRepo) error {
	sqlStr := "INSERT INTO PhoneBook(CountryCode, CountryName, PhoneCode) VALUES ($1,$2,$3)"
	country := data.CountryName
	phone := data.PhoneCode
	tx, err:=db.Db.Begin()
	for key, val := range country {
		_,err = tx.Exec(sqlStr,key, fmt.Sprintf("%s",val), fmt.Sprintf("%s",phone[key]))
		if err != nil {
			return err
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

func (d Postgresrepo) Select (country string ) ([]schema.PhoneEntity,error) {
	rows, err := d.Db.Query("select * from PhoneBook where upper(CountryName)=$1 LIMIT 1", strings.ToUpper(country))
	defer rows.Close()
	if err != nil {
		return nil,err
	}
	var products []schema.PhoneEntity
	for rows.Next() {
		p := schema.PhoneEntity{}
		err := rows.Scan(&p.CountryCode, &p.CountryName, &p.PhoneCode)
		if err != nil {
			return nil,err
		}
		products = append(products, p)
	}

	return products,nil
}