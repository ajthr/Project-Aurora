package config

import (
	"strconv"

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

func (c *JWTConfig) CreateToken(userId int) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": strconv.Itoa(userId),
		})
	token, err := t.SignedString(c.SecretKey)

	if err != nil {
		return "", err
	}

	return token, nil
}
