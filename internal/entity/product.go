package entity

import (
	"errors"
	"github.com/IcaroTARique/RamenGoRV/pkg"
)

type ProductType string

var (
	ErrIDIsRequired            = errors.New("ID is required")
	ErrImageInactiveIsRequired = errors.New("imageInactive is required")
	ErrImageActiveIsRequired   = errors.New("imageActive is required")
	ErrNameIsRequired          = errors.New("name is required")
	ErrDescriptionIsRequired   = errors.New("description is required")
	ErrPriceIsRequired         = errors.New("price is required")
	ErrorInvalidType           = errors.New("invalid type choosen is not valid")
	TypeBroth                  = "broth"
	TypeProtein                = "protein"
)

type Product struct {
	Id            int64   `json:"id"`
	ImageInactive string  `json:"imageInactive"`
	ImageActive   string  `json:"imageActive"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	TypeOfProduct string  `json:"type"`
	Price         float64 `json:"price"`
}

func NewProduct(id int, imageInactive, imageActive, name, description, typeOf string, price float64) *Product {
	return &Product{
		Id:            pkg.ParseIdToInt64(id),
		ImageInactive: imageInactive,
		ImageActive:   imageActive,
		Name:          name,
		Description:   description,
		TypeOfProduct: typeOf,
		Price:         price,
	}
}

func (p *Product) Validate() error {

	if p.ImageInactive == "" {
		return ErrImageInactiveIsRequired

	}

	if p.ImageActive == "" {
		return ErrImageActiveIsRequired

	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if p.Description == "" {
		return ErrDescriptionIsRequired
	}

	if p.Price <= 0 {
		return ErrPriceIsRequired
	}

	if err := p.ValidateType(); err != nil {
		return err
	}

	return nil
}

func (p *Product) ValidateType() error {
	if p.TypeOfProduct != TypeBroth && p.TypeOfProduct != TypeProtein {
		return ErrorInvalidType
	}
	return nil
}
