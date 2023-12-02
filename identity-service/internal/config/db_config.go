package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	DB *sql.DB
}

func NewDBConfig(host, port, user, password, dbname string) (*DBConfig, error) {

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname)

	db, err := sql.Open("postgres", connection)

	return &DBConfig{
		DB: db,
	}, err
}
