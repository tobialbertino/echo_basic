package fruitApi

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s Service) GetAllData(_ context.Context) (res []DataFruit, err error) {
	var netClient = &http.Client{}

	req, err := http.NewRequest(http.MethodGet, BaseUrl+UrlGetAllData, bytes.NewBufferString(""))
	if err != nil {
		return
	}

	resp, err := netClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return
	}

	return
}
