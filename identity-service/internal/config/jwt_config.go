package config

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	SecretKey   []byte
	TokenParser *jwt.Parser
}

func NewJWTConfig() *JWTConfig {
	return &JWTConfig{
		SecretKey:   []byte(os.Getenv("JWT_SECRET")),
		TokenParser: jwt.NewParser(),
	}
}
