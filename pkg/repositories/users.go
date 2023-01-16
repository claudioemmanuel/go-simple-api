package repositories

import (
	"database/sql"
	"fmt"

	"devbook-api/pkg/models"
)

// UserRepository is a struct that represents a user repository
type UserRepository struct {
	db *sql.DB
}

func GetUsers() ([]models.User, error) {
	return []models.User{}, nil
}

// CreateUser creates a new user
func (u UserRepository) CreateUser() error {
	return fmt.Errorf("Not implemented")
}
