package entity

import (
	"testing"

	"github.com/eralves01/api-rest-jwt/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := entity.NewUser("Eric Alves", "alves@example.com", "password")
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Eric Alves", user.Name)
	assert.Equal(t, "alves@example.com", user.Email)
}

func TestValidatePassword(t *testing.T) {
	user, _ := entity.NewUser("Eric Alves", "alves@example.com", "password")
	assert.NotEqual(t, user.Password, "password")
	assert.True(t, user.ValidatePassword("password"))
	assert.False(t, user.ValidatePassword("password123"))
}
