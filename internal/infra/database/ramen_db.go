package database

import (
	"github.com/IcaroTARique/RamenGoRV/internal/dto"
	api "github.com/IcaroTARique/RamenGoRV/internal/infra/client/api_id"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) GetBroths() ([]dto.Broth, error) {
	var broths []dto.Broth
	err := p.DB.Find(&broths).Error
	return broths, err
}

func (p *Product) GetProteins() ([]dto.Protein, error) {
	var proteins []dto.Protein
	err := p.DB.Find(&proteins).Error
	return proteins, err
}

func (p *Product) CreateOrder(order dto.OrderRequest) (*dto.OrderResponse, error) {

	var orderResponse dto.OrderResponse
	err := p.DB.Create(order).Error
	if err != nil {
		return nil, err
	}

	clientApi := api.NewClientIdConnection()
	clientId, _ := clientApi.GetProductId()
	orderResponse.Id = clientId.OrderId
	return &orderResponse, nil
}
