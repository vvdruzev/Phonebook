package data

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"log"
	"time"
	"fmt"
	"Phonebook/logger"
)



type DataRepo struct {
	PhoneCode map[string]interface{}
	CountryName map[string]interface{}
}

func NewDataRepo() *DataRepo  {
	return &DataRepo{}
}

type ResourceError struct {
	URL string
	Err error
}


func (re *ResourceError) Error() string {
	return fmt.Sprintf(
		"Resource error: URL: %s, err: %v",
		re.URL,
		re.Err,
	)
}
var client http.Client

func SetClient(proxy string)  {
	if proxy =="" {
		client = http.Client{}
	}else {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			log.Println(err)
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}

		client = http.Client{
			Transport: transport,
			Timeout:   time.Second * 1,
		}
	}
}

func (d *DataRepo)  GetPhoneCode() (error) {
	url:="http://country.io/phone.json"
	logger.Info("Getting data from ", url)
	resp, err := client.Get(url)
	if err != nil {
		return &ResourceError{URL: url, Err: err}
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return &ResourceError{URL: url, Err: err}
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
	logger.Info("Getting data from ", url)
	resp, err := client.Get(url)
	if err != nil {
		return &ResourceError{URL: url, Err: err}
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return &ResourceError{URL: url, Err: err}
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

