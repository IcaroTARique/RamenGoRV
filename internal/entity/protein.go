package entity

import (
	"github.com/IcaroTARique/RamenGoRV/pkg"
)

type Protein struct {
	Id            int64  `json:"id"`
	ImageInactive string `json:"imageInactive"`
	ImageActive   string `json:"imageActive"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         int    `json:"price"`
}

func (*Protein) TableName() string {
	return "protein"
}

func NewProtein(id int, imageInactive, imageActive, name, description string, price int) *Protein {
	return &Protein{
		Id:            pkg.ParseIdToInt64(id),
		ImageInactive: imageInactive,
		ImageActive:   imageActive,
		Name:          name,
		Description:   description,
		Price:         price,
	}
}

func (p *Protein) Validate() error {

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
