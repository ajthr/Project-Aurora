package database

import (
	"database/sql"
)

type AuthStore struct {
	db *sql.DB
}

func NewAuthStore(conn *sql.DB) *AuthStore {
	return &AuthStore{
		db: conn,
	}
}
