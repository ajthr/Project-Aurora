package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	DB *sql.DB
}

func initDatabase(db *sql.DB) {
	initScript := `
		CREATE TABLE IF NOT EXISTS users (
			id					serial			PRIMARY KEY,
			name				varchar(128)	NOT NULL 		DEFAULT '',
			email				varchar(255)	UNIQUE,
			is_signup_complete	boolean,
			user_created_at		timestamp
		);

		CREATE TABLE IF NOT EXISTS otp (
			id					serial			PRIMARY KEY,
			email				varchar(255)	UNIQUE,
			otp					varchar(16),
			expiration			timestamp
		);
	`

	_, err := db.Exec(initScript)

	if err != nil {
		log.Panic("WARN cannot run init script", err)
	}
}

func NewDBConfig(host, port, user, password, dbname string) (*DBConfig, error) {

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname)

	db, err := sql.Open("postgres", connection)

	// initialise database with required tables on successful connection
	if err == nil {
		initDatabase(db)
	}

	return &DBConfig{
		DB: db,
	}, err
}
