package database

import "database/sql"

type AccountStore struct {
	db *sql.DB
}

func NewAccountStore(conn *sql.DB) *AccountStore {
	return &AccountStore{
		db: conn,
	}
}
