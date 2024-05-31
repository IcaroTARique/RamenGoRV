package api

import (
	"encoding/json"
	"github.com/IcaroTARique/RamenGoRV/internal/dto"
	"net/http"
)

var (
	acquireIdPath = "/orders/generate-id"
)

type ClientIdConnection struct {
	BaseURL     string
	APIKey      string
	APIKeyValue string
}

func NewClientIdConnection() *ClientIdConnection {
	return &ClientIdConnection{
		BaseURL:     "https://api.tech.redventures.com.br",
		APIKey:      "x-api-key",
		APIKeyValue: "ZtVdh8XQ2U8pWI2gmZ7f796Vh8GllXoN7mr0djNf",
	}
}

func (c *ClientIdConnection) GetProductId() (*dto.OrderId, error) {

	URL := c.BaseURL + acquireIdPath
	req, err := http.NewRequest("POST", URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set(c.APIKey, c.APIKeyValue)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var OrderId dto.OrderId
	err = json.NewDecoder(res.Body).Decode(&OrderId)
	if err != nil {
		return nil, err
	}
	return &OrderId, nil
}
