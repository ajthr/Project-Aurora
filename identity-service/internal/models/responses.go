package models

// jwt token response
type JWTTokenResponse struct {
	Token            string
	IsSignupComplete bool
}
