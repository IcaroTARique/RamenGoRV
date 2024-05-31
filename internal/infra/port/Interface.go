package port

import "github.com/IcaroTARique/RamenGoRV/internal/dto"

type ProductInterface interface {
	GetBroths() ([]dto.Broth, error)
	GetProteins() ([]dto.Protein, error)
	CreateOrder(order dto.OrderRequest) (*dto.OrderResponse, error)
}
