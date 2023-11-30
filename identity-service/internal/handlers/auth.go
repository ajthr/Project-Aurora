package handlers

import (
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GoogleSignIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
