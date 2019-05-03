package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DataRepo struct {
	PhoneCode map[string]interface{}
	CountryName map[string]interface{}
}

func NewDataRepo() *DataRepo  {
	return &DataRepo{}
}
func (d *DataRepo)  GetPhoneCode() (error) {
	url:="http://country.io/phone.json"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	d.PhoneCode = data
	return  nil
}

func (d *DataRepo) GetCountryName()  error  {
	url:="http://country.io/names.json"
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return err
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	d.CountryName = data
	return  nil
}

