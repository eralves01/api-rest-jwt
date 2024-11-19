package database

import (
	"testing"

	"github.com/eralves01/api-rest-jwt/internal/entity"
	"github.com/eralves01/api-rest-jwt/internal/infra/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=api-rest-jwt port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection error!")
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Eric Alves", "alves@example.com", "password")
	userDB := database.NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, "Eric Alves", userFound.Name)
	assert.Equal(t, "alves@example.com", userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	dsn := "host=localhost user=postgres password=postgres dbname=api-rest-jwt port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database connection error!")
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Eric Alves", "alves@example.com", "password")
	userDB := database.NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, "Eric Alves", userFound.Name)
	assert.Equal(t, "alves@example.com", userFound.Email)
	assert.NotNil(t, userFound.Password)
}
