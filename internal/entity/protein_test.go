package entity

import (
	"github.com/IcaroTARique/RamenGoRV/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProtein(t *testing.T) {
	protein := NewProtein(1, "imageInactive", "imageActive", "name", "protein", 1)
	assert.Equal(t, pkg.ParseIdToInt64(1), protein.Id)
	assert.NotEmpty(t, protein.Id)
	assert.Equal(t, "imageInactive", protein.ImageInactive)
	assert.Equal(t, "imageActive", protein.ImageActive)
	assert.Equal(t, "name", protein.Name)
	assert.Equal(t, 1, protein.Price)
	assert.NotZero(t, protein.Price)
}

func TestProteinWhenIdIsZero(t *testing.T) {
	protein := NewProtein(0, "imageInactive", "imageActive", "name", "protein", 1)
	assert.NotNil(t, protein)
	assert.Error(t, ErrIDIsRequired, protein.Validate())
}

func TestProteinBrothWhenIdIsNegative(t *testing.T) {
	protein := NewProtein(-1, "imageInactive", "imageActive", "name", "protein", 1)
	assert.NotNil(t, protein)
	assert.Error(t, ErrIDIsRequired, protein.Validate())
}

func TestProteinBrothWhenIdIsMaximumInt64(t *testing.T) {
	protein := NewProtein(9223372036854775807, "imageInactive", "imageActive", "name", "protein", 1)
	assert.NotNil(t, protein)
	assert.NotEmpty(t, protein.Id)
	assert.Equal(t, int64(9223372036854775807), protein.Id)
}

func TestProteinImageInactiveIsRequired(t *testing.T) {
	protein := NewProtein(1, "", "imageActive", "name", "protein", 1)
	assert.NotNil(t, protein)
	assert.Equal(t, ErrImageInactiveIsRequired, protein.Validate())
}

func TestProteinImageActiveIsRequired(t *testing.T) {
	protein := NewProtein(1, "imageInactive", "", "name", "protein", 1)
	assert.NotNil(t, protein)
	assert.Equal(t, ErrImageActiveIsRequired, protein.Validate())
}

func TestProteinNameIsRequired(t *testing.T) {
	protein := NewProtein(1, "imageInactive", "imageActive", "", "protein", 1)
	assert.NotNil(t, protein)
	assert.Equal(t, ErrNameIsRequired, protein.Validate())
}

func TestProteinDescriptionIsRequired(t *testing.T) {
	protein := NewProtein(1, "imageInactive", "imageActive", "name", "", 1)
	assert.NotNil(t, protein)
	assert.Equal(t, ErrDescriptionIsRequired, protein.Validate())
}

func TestProteinPriceIsZero(t *testing.T) {
	protein := NewProtein(1, "imageInactive", "imageActive", "name", "protein", 0)
	assert.NotNil(t, protein)
	assert.Equal(t, ErrPriceIsRequired, protein.Validate())
}

func TestProteinPriceIsNegative(t *testing.T) {
	protein := NewProtein(1, "imageInactive", "imageActive", "name", "protein", -1)
	assert.NotNil(t, protein)
	assert.Equal(t, ErrPriceIsRequired, protein.Validate())
}

func TestProteinWrongType(t *testing.T) {
	protein := NewProtein(1, "imageInactive", "imageActive", "name", "protein", 1)
	assert.NotNil(t, protein)
	assert.Error(t, ErrorInvalidType, protein.Validate())
}
