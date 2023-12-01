package middlewares

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"identity-service/internal/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

// utility functions for testing

// create a valid jwt token
func createJwtToken(userId string) string {
	key := os.Getenv("JWT_SECRET")
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userId,
		})
	token, _ := t.SignedString([]byte(key))
	return token
}

// create an invalid jwt token
func createInvalidJwtToken(userId string) string {
	key := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	t := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userId,
		})
	token, _ := t.SignedString([]byte(key))
	return token
}

// create an invalid jwt token
func createJwtTokenWithoutClaims(userId string) string {
	key := os.Getenv("JWT_SECRET")
	t := jwt.New(jwt.SigningMethodHS256)
	token, _ := t.SignedString([]byte(key))
	return token
}

// function to test authenticator middleware
func AuthenticatorMiddlewareResponse(userId string, token string) *http.Response {
	// create config for authenticator
	config := config.NewJWTConfig()

	// mock handler fn that returns 200 once the request reaches it
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Add("Authorization", "Bearer "+token)

	next := Authenticator(*config)(handler)
	next.ServeHTTP(recorder, request)

	return recorder.Result()
}

func TestAuthenticatorMiddlewareWithValidToken(t *testing.T) {

	var userIds = []string{"12345", "54321"}

	for _, userId := range userIds {
		token := createJwtToken(userId)
		response := AuthenticatorMiddlewareResponse(userId, token)

		// check statuscode
		assert.Equal(t, 200, response.StatusCode)

		// check if the header is set
		assert.NotEmpty(t, response.Header.Get("X-Authenticated-User-ID"))

		// check if the correct header is being set
		assert.Equal(t, userId, response.Header.Get("X-Authenticated-User-ID"))
	}

}

func TestAuthenticatorMiddlewareWithInValidToken(t *testing.T) {

	var userIds = []string{"12345", "54321"}

	for _, userId := range userIds {
		token := createInvalidJwtToken(userId)
		response := AuthenticatorMiddlewareResponse(userId, token)

		// check statuscode
		assert.Equal(t, 401, response.StatusCode)

		// check if the header is empty
		assert.Empty(t, response.Header.Get("X-Authenticated-User-ID"))
	}

}

func TestAuthenticatorMiddlewareWithValidTokenWithoutClaims(t *testing.T) {

	var userIds = []string{"12345", "54321"}

	for _, userId := range userIds {
		token := createJwtTokenWithoutClaims(userId)
		response := AuthenticatorMiddlewareResponse(userId, token)

		// check statuscode
		assert.Equal(t, 401, response.StatusCode)

		// check if the header is empty
		assert.Empty(t, response.Header.Get("X-Authenticated-User-ID"))
	}

}
