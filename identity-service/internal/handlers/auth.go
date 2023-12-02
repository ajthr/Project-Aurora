package handlers

import (
	"database/sql"
	"net/http"

	"identity-service/internal/database"
)

type AuthHandler struct {
	store *database.AuthStore
}

func NewAuthHandler(conn *sql.DB) *AuthHandler {
	authStore := database.NewAuthStore(conn)
	return &AuthHandler{
		store: authStore,
	}
}

// function to signin and create user if not exists
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// function to signin and signup with google auth
func (h *AuthHandler) GoogleSignIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// function to validate otp for signin and signup
func (h *AuthHandler) ValidateOtp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
