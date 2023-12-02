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

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) GoogleSignIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
