package entity

import (
	"github.com/IcaroTARique/RamenGoRV/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBroth(t *testing.T) {
	broth := NewBroth(1, "imageInactive", "imageActive", "name", "protein", 1)
	assert.Equal(t, pkg.ParseIdToInt64(1), broth.Id)
	assert.NotEmpty(t, broth.Id)
	assert.Equal(t, "imageInactive", broth.ImageInactive)
	assert.Equal(t, "imageActive", broth.ImageActive)
	assert.Equal(t, "name", broth.Name)
	assert.Equal(t, 1, broth.Price)
	assert.NotZero(t, broth.Price)
}

func TestBrothWhenIdIsZero(t *testing.T) {
	broth := NewBroth(0, "imageInactive", "imageActive", "name", "broth", 1)
	assert.NotNil(t, broth)
	assert.Error(t, ErrIDIsRequired, broth.Validate())
}

func TestBrothWhenIdIsNegative(t *testing.T) {
	broth := NewBroth(-1, "imageInactive", "imageActive", "name", "broth", 1)
	assert.NotNil(t, broth)
	assert.Error(t, ErrIDIsRequired, broth.Validate())
}

func TestBrothWhenIdIsMaximumInt64(t *testing.T) {
	broth := NewBroth(9223372036854775807, "imageInactive", "imageActive", "name", "broth", 1)
	assert.NotNil(t, broth)
	assert.NotEmpty(t, broth.Id)
	assert.Equal(t, int64(9223372036854775807), broth.Id)
}

func TestImageInactiveIsRequired(t *testing.T) {
	broth := NewBroth(1, "", "imageActive", "name", "broth", 1)
	assert.NotNil(t, broth)
	assert.Equal(t, ErrImageInactiveIsRequired, broth.Validate())
}

func TestImageActiveIsRequired(t *testing.T) {
	broth := NewBroth(1, "imageInactive", "", "name", "broth", 1)
	assert.NotNil(t, broth)
	assert.Equal(t, ErrImageActiveIsRequired, broth.Validate())
}

func TestNameIsRequired(t *testing.T) {
	broth := NewBroth(1, "imageInactive", "imageActive", "", "broth", 1)
	assert.NotNil(t, broth)
	assert.Equal(t, ErrNameIsRequired, broth.Validate())
}

func TestDescriptionIsRequired(t *testing.T) {
	broth := NewBroth(1, "imageInactive", "imageActive", "name", "", 1)
	assert.NotNil(t, broth)
	assert.Equal(t, ErrDescriptionIsRequired, broth.Validate())
}

func TestPriceIsZero(t *testing.T) {
	broth := NewBroth(1, "imageInactive", "imageActive", "name", "broth", 0)
	assert.NotNil(t, broth)
	assert.Equal(t, ErrPriceIsRequired, broth.Validate())
}

func TestPriceIsNegative(t *testing.T) {
	broth := NewBroth(1, "imageInactive", "imageActive", "name", "broth", -1)
	assert.NotNil(t, broth)
	assert.Equal(t, ErrPriceIsRequired, broth.Validate())
}

func TestWrongType(t *testing.T) {
	broth := NewBroth(1, "imageInactive", "imageActive", "name", "broth", 1)
	assert.NotNil(t, broth)
	assert.Error(t, ErrorInvalidType, broth.Validate())
}
