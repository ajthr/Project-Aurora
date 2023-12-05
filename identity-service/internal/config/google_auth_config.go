package config

import (
	"context"

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

func NewGoogleAuthClient(googleClientId string) *GoogleAuthClient {
	return &GoogleAuthClient{
		ClientId: googleClientId,
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
