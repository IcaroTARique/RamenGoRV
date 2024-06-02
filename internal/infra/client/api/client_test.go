package api

import (
	"fmt"
	"github.com/IcaroTARique/RamenGoRV/internal/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBroths(t *testing.T) {
	clientBroth := NewClient()
	broths, err := clientBroth.GetBroths()
	fmt.Println(broths)
	assert.Nil(t, err)
	assert.NotNil(t, broths)
}

func TestGetProteins(t *testing.T) {
	clientProtein := NewClient()
	proteins, err := clientProtein.GetProtein()
	fmt.Println(proteins)
	assert.Nil(t, err)
	assert.NotNil(t, proteins)
}

func TestCreateOrder(t *testing.T) {
	clientCreateOrder := NewClient()
	order, err := clientCreateOrder.CreateOrder(
		dto.OrderRequest{BrothId: "1", ProteinId: "1"},
	)
	assert.Nil(t, err)
	assert.NotNil(t, order)
	fmt.Println(order)
}
