package api

import (
	"bytes"
	"encoding/json"
	"github.com/IcaroTARique/RamenGoRV/internal/dto"
	"net/http"
)

var (
	brothPath       = "/broths"
	proteinPath     = "/proteins"
	createOrderPath = "/orders"
)

type ClientConnection struct {
	BaseURL     string
	APIKey      string
	APIKeyValue string
}

func NewClient() *ClientConnection {
	return &ClientConnection{
		BaseURL:     "https://api.tech.redventures.com.br",
		APIKey:      "x-api-key",
		APIKeyValue: "ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf",
	}
}

func (c *ClientConnection) GetBroths() ([]dto.Broth, error) {

	URL := c.BaseURL + brothPath
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(c.APIKey, c.APIKeyValue)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var brothValues []dto.Broth
	err = json.NewDecoder(res.Body).Decode(&brothValues)
	if err != nil {
		return nil, err
	}
	return brothValues, nil
}

func (c *ClientConnection) GetProteins() ([]dto.Protein, error) {

	URL := c.BaseURL + proteinPath
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(c.APIKey, c.APIKeyValue)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var proteinesValue []dto.Protein
	err = json.NewDecoder(res.Body).Decode(&proteinesValue)
	if err != nil {
		return nil, err
	}
	return proteinesValue, nil
}

func (c *ClientConnection) CreateOrder(orderRequest dto.OrderRequest) (*dto.OrderResponse, error) {
	URL := c.BaseURL + createOrderPath
	marshalledBody, err := json.Marshal(orderRequest)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", URL, bytes.NewReader(marshalledBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set(c.APIKey, c.APIKeyValue)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	var orderResponse dto.OrderResponse
	err = json.NewDecoder(res.Body).Decode(&orderResponse)
	if err != nil {
		return nil, err
	}
	return &orderResponse, nil
}
