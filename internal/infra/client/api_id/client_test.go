package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProductId(t *testing.T) {
	c := NewClientIdConnection()
	id, err := c.GetProductId()
	assert.NotNil(t, id)
	assert.Nil(t, err)
}
