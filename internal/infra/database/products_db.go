package database

import (
	"fmt"
	"github.com/IcaroTARique/RamenGoRV/internal/dto"
	"github.com/IcaroTARique/RamenGoRV/internal/entity"
	api "github.com/IcaroTARique/RamenGoRV/internal/infra/client/api_id"
	"github.com/IcaroTARique/RamenGoRV/pkg"
	"gorm.io/gorm"
)

type ProductDB struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *ProductDB {
	return &ProductDB{DB: db}
}

func (p *ProductDB) GetBroths() ([]dto.Broth, error) {
	var broths []entity.Broth
	err := p.DB.Find(&broths).Error
	fmt.Println("BROTHS", broths)
	if err != nil {
		return nil, err
	}

	var brothsDto []dto.Broth
	for _, broth := range broths {
		brothDto := dto.Broth{
			Id:            pkg.ParseIdToString(broth.Id),
			ImageInactive: broth.ImageInactive,
			ImageActive:   broth.ImageActive,
			Name:          broth.Name,
			Description:   broth.Description,
			Price:         broth.Price,
		}
		brothsDto = append(brothsDto, brothDto)
	}
	return brothsDto, nil
}

func (p *ProductDB) GetProtein() ([]dto.Protein, error) {
	var protein []entity.Protein
	err := p.DB.Find(&protein).Error
	if err != nil {
		return nil, err
	}

	var proteinsDto []dto.Protein
	for _, protein := range protein {
		proteinDto := dto.Protein{
			Id:            pkg.ParseIdToString(protein.Id),
			ImageInactive: protein.ImageInactive,
			ImageActive:   protein.ImageActive,
			Name:          protein.Name,
			Description:   protein.Description,
			Price:         protein.Price,
		}
		proteinsDto = append(proteinsDto, proteinDto)
	}

	return proteinsDto, err
}

func (p *ProductDB) CreateOrder(orderRequest dto.OrderRequest) (*dto.OrderResponse, error) {

	orderEntity := entity.NewOrder(orderRequest.BrothId, orderRequest.ProteinId)

	err := p.DB.Create(&orderEntity).Error
	if err != nil {
		return nil, err
	}

	response, err := p.GetResponse(orderRequest.BrothId, orderRequest.ProteinId)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (p *ProductDB) GetResponse(brothId, proteinId string) (*dto.OrderResponse, error) {
	var answer entity.Answer
	err := p.DB.Find(&answer).Where("broth_id = ? AND protein_id = ?", brothId, proteinId).Error
	if err != nil {
		return nil, err
	}

	clientApi := api.NewClientIdConnection()
	clientId, _ := clientApi.GetProductId()
	return &dto.OrderResponse{
		Id:          clientId.OrderId,
		Description: answer.Description,
		Image:       answer.Image,
	}, nil
}
