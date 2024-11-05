package entity

import (
	"errors"
	"time"

	"github.com/eralves01/api-rest-jwt/pkg/entity"
)

var (
	ErrorIDIsRequired    = errors.New("id is required")
	ErrorInvalidID       = errors.New("invalid id")
	ErrorNameIsRequired  = errors.New("name is required")
	ErrorPriceIsRequired = errors.New("price is required")
	ErrorInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := product.IsValid(); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) IsValid() error {
	if p.ID.String() == "" {
		return ErrorIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrorInvalidID
	}
	if p.Name == "" {
		return ErrorNameIsRequired
	}
	if p.Price == 0 {
		return ErrorPriceIsRequired
	}
	if p.Price < 0 {
		return ErrorInvalidPrice
	}
	return nil
}
