package entity

import (
	"testing"

	"github.com/eralves01/api-rest-jwt/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := entity.NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Product 1", product.Name)
	assert.Equal(t, 10, product.Price)

	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreatedAt)
	assert.NotEmpty(t, product.UpdatedAt)
}

func TestIsValid(t *testing.T) {
	product, _ := entity.NewProduct("Product 1", 10)
	assert.Nil(t, product.IsValid())
}

func TestProductWhenNameIsEmpty(t *testing.T) {
	p, err := entity.NewProduct("", 10)
	assert.Nil(t, p)
	assert.Equal(t, entity.ErrorNameIsRequired, err)
}

func TestProductWhenPriceIsZero(t *testing.T) {
	p, err := entity.NewProduct("Product 1", 0)
	assert.Nil(t, p)
	assert.Equal(t, entity.ErrorPriceIsRequired, err)
}

func TestProductWhenPriceIsNegative(t *testing.T) {
	p, err := entity.NewProduct("Product 1", -10)
	assert.Nil(t, p)
	assert.Equal(t, entity.ErrorInvalidPrice, err)
}
