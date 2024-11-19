package database

import (
	"github.com/eralves01/api-rest-jwt/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
}
