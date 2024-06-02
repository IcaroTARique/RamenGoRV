package entity

import (
	"errors"
	"github.com/IcaroTARique/RamenGoRV/pkg"
)

var (
	ErrIDIsRequired            = errors.New("ID is required")
	ErrImageInactiveIsRequired = errors.New("imageInactive is required")
	ErrImageActiveIsRequired   = errors.New("imageActive is required")
	ErrNameIsRequired          = errors.New("name is required")
	ErrDescriptionIsRequired   = errors.New("description is required")
	ErrPriceIsRequired         = errors.New("price is required")
	ErrorInvalidType           = errors.New("invalid type choosen is not valid")
)

type Broth struct {
	Id            int64  `json:"id"`
	ImageInactive string `json:"imageInactive"`
	ImageActive   string `json:"imageActive"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
}

func (*Broth) TableName() string {
	return "broth"
}

func NewBroth(id int, imageInactive, imageActive, name, description string, price int) *Broth {
	return &Broth{
		Id:            pkg.ParseIdToInt64(id),
		ImageInactive: imageInactive,
		ImageActive:   imageActive,
		Name:          name,
		Description:   description,
		Price:         price,
	}
}

func (p *Broth) Validate() error {

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

	return nil
}
