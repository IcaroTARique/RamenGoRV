package entity

import (
	"github.com/IcaroTARique/RamenGoRV/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	product := NewProduct(1, "imageInactive", "imageActive", "name", "protein", "protein", 1.0)
	assert.Equal(t, pkg.ParseIdToInt64(1), product.Id)
	assert.NotEmpty(t, product.Id)
	assert.Equal(t, "imageInactive", product.ImageInactive)
	assert.Equal(t, "imageActive", product.ImageActive)
	assert.Equal(t, "name", product.Name)
	assert.Equal(t, 1.0, product.Price)
	assert.NotZero(t, product.Price)
}

func TestProductWhenIdIsZero(t *testing.T) {
	product := NewProduct(0, "imageInactive", "imageActive", "name", "broth", "broth", 1.0)
	assert.NotNil(t, product)
	assert.Error(t, ErrIDIsRequired, product.Validate())
}

func TestProductWhenIdIsNegative(t *testing.T) {
	product := NewProduct(-1, "imageInactive", "imageActive", "name", "broth", "broth", 1.0)
	assert.NotNil(t, product)
	assert.Error(t, ErrIDIsRequired, product.Validate())
}

func TestProductWhenIdIsMaximumInt64(t *testing.T) {
	product := NewProduct(9223372036854775807, "imageInactive", "imageActive", "name", "protein", "protein", 1.0)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.Id)
	assert.Equal(t, int64(9223372036854775807), product.Id)
}

func TestImageInactiveIsRequired(t *testing.T) {
	product := NewProduct(1, "", "imageActive", "name", "protein", "protein", 1.0)
	assert.NotNil(t, product)
	assert.Equal(t, ErrImageInactiveIsRequired, product.Validate())
}

func TestImageActiveIsRequired(t *testing.T) {
	product := NewProduct(1, "imageInactive", "", "name", "protein", "protein", 1.0)
	assert.NotNil(t, product)
	assert.Equal(t, ErrImageActiveIsRequired, product.Validate())
}

func TestNameIsRequired(t *testing.T) {
	product := NewProduct(1, "imageInactive", "imageActive", "", "protein", "protein", 1.0)
	assert.NotNil(t, product)
	assert.Equal(t, ErrNameIsRequired, product.Validate())
}

func TestDescriptionIsRequired(t *testing.T) {
	product := NewProduct(1, "imageInactive", "imageActive", "name", "", "protein", 1.0)
	assert.NotNil(t, product)
	assert.Equal(t, ErrDescriptionIsRequired, product.Validate())
}

func TestPriceIsZero(t *testing.T) {
	product := NewProduct(1, "imageInactive", "imageActive", "name", "protein", "protein", 0)
	assert.NotNil(t, product)
	assert.Equal(t, ErrPriceIsRequired, product.Validate())
}

func TestPriceIsNegative(t *testing.T) {
	product := NewProduct(1, "imageInactive", "imageActive", "name", "protein", "protein", -1.0)
	assert.NotNil(t, product)
	assert.Equal(t, ErrPriceIsRequired, product.Validate())
}

func TestWrongType(t *testing.T) {
	product := NewProduct(1, "imageInactive", "imageActive", "name", "protein", "invalid", 1.0)
	assert.NotNil(t, product)
	assert.Error(t, ErrorInvalidType, product.Validate())
}
