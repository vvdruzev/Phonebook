package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	COUNTRYNAME = iota
	PHONECODE
)

func GetData(f int) (map[string]interface{},error) {
	url:=""
	if f==COUNTRYNAME {
		url="http://country.io/names.json"
	}
	if f==PHONECODE {
		url="http://country.io/phone.json"
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil,err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil,err
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil,err
	}

	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil,err
	}
	return data, nil

}



