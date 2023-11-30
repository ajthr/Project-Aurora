package handlers

import (
	"net/http"
)

func ChangeMail(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
