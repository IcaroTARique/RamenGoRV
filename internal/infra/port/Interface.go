package port

import "github.com/IcaroTARique/RamenGoRV/internal/dto"

type ProductInterface interface {
	GetBroths() ([]dto.Broth, error)
	GetProtein() ([]dto.Protein, error)
	CreateOrder(order dto.OrderRequest) (*dto.OrderResponse, error)
}
