package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type HttpClientConfig struct {
	Hostname string
}

func New(hostname string) HttpClientConfig {

	config := HttpClientConfig{
		Hostname: hostname,
	}

	return config

}

func call(requestURL string, object interface{}) error {
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(object)
	if err != nil {
		log.Println(err)
		return err
	}

	return err
}
