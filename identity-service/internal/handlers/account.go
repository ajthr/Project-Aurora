package handlers

import (
	"database/sql"
	"identity-service/internal/database"
	"net/http"
)

type AccountHandler struct {
	store *database.AccountStore
}

func NewAccountHandler(conn *sql.DB) *AccountHandler {
	accountStore := database.NewAccountStore(conn)
	return &AccountHandler{
		store: accountStore,
	}
}

func (h *AccountHandler) ChangeMail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *AccountHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
