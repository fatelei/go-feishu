package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Transport struct {
	Endpoint string
}

func (p *Transport) Get(path string, param *map[string]string, accessToken string) ([]byte, error) {
	baseUrl := fmt.Sprintf("%s%s", p.Endpoint, path)
	query := url.Values{}
	for k, v := range *param {
		query.Add(k, v)
	}
	queryParam := query.Encode()
	log.Printf("%s?%s\n", baseUrl, queryParam)
	url := fmt.Sprintf("%s?%s", baseUrl, queryParam)
	client := http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (p *Transport) Post(path string, param *map[string]interface{}, accessToken string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", p.Endpoint, path)
	log.Printf("%s\n", url)
	client := http.Client{}
	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (p *Transport) PostJson(path string, param map[string]interface{}) ([]byte, error) {
	baseUrl := fmt.Sprintf("%s%s", p.Endpoint, path)
	request, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	log.Printf("%s %v\n", baseUrl, param)
	resp, err := http.Post(baseUrl, "application/json", bytes.NewBuffer(request))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}