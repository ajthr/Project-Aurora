package config

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTConfig struct {
	SecretKey   []byte
	TokenParser *jwt.Parser
}

func NewJWTConfig(secretKey string) *JWTConfig {
	return &JWTConfig{
		SecretKey:   []byte(secretKey),
		TokenParser: jwt.NewParser(),
	}
}
