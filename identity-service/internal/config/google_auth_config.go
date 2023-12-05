package config

import (
	"context"
	"os"

	"google.golang.org/api/idtoken"
)

type GoogleClaims struct {
	Email string
	Name  string
}

type GoogleAuthClient struct {
	ClientId string
	Claims   *GoogleClaims
}

func NewGoogleAuthClient() *GoogleAuthClient {
	return &GoogleAuthClient{
		ClientId: os.Getenv("GOOGLE_CLIENT_ID"),
		Claims:   &GoogleClaims{},
	}
}

func (g *GoogleAuthClient) VerifyGoogleToken(token string) error {
	payload, err := idtoken.Validate(context.Background(), token, g.ClientId)

	if err != nil {
		return err
	}

	g.Claims = &GoogleClaims{
		Email: payload.Claims["email"].(string),
		Name:  payload.Claims["given_name"].(string),
	}

	return nil
}
