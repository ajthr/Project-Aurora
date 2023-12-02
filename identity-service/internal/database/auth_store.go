package database

import (
	"database/sql"
	"identity-service/internal/models"
)

type AuthStore struct {
	db *sql.DB
}

func NewAuthStore(conn *sql.DB) *AuthStore {
	return &AuthStore{
		db: conn,
	}
}

func (s *AuthStore) FindUserById() (*models.User, error) {
	return &models.User{}, nil
}

func (s *AuthStore) FindUserByEmail() (*models.User, error) {
	return &models.User{}, nil
}

func (s *AuthStore) CreateUser(user models.User) error {
	return nil
}

func (s *AuthStore) FindOtpByUserId(id int64) {

}

func (s *AuthStore) CreateOtp() error {
	return nil
}

func (s *AuthStore) UpdateOtp() error {
	return nil
}
