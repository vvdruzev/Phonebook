package db

import (
	"Phonebook/data"
	"Phonebook/schema"
)

type Repository interface {
	Close()
	Reload (repo data.DataRepo) error
	Select(country string ) (schema.ResponseCode,error)
	Insert(repo data.DataRepo) error

}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Close()
}

func Reload(repo data.DataRepo) error  {
	return impl.Reload(repo)
}

func Select(country string ) (schema.ResponseCode,error)  {
	return impl.Select(country)
}

func Insert(repo data.DataRepo) error {
	return impl.Insert(repo)
}