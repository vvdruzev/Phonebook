package db

import (
	"Phonebook/data"
	"Phonebook/schema"
)

type Repository interface {
	Create() error
	Reload (repo data.DataRepo) error
	Select(country string ) ([]schema.PhoneEntity,error)
	Insert(repo data.DataRepo) error

}

var impl Repository

func SetRepository(repository Repository) {
	impl = repository
}

func Close() {
	impl.Create()
}

func Reload(repo data.DataRepo) error  {
	return impl.Reload(repo)
}

func Select(country string ) ([]schema.PhoneEntity,error)  {
	return impl.Select(country)
}

func Insert(repo data.DataRepo) error {
	return impl.Insert(repo)
}