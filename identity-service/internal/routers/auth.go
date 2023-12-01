package routers

import (
	"identity-service/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func AuthRouter() *chi.Mux {
	router := chi.NewRouter()

	// route handlers
	router.Post("/signin", handlers.SignIn)
	router.Post("/signup", handlers.SignUp)
	router.Post("/google-signin", handlers.GoogleSignIn)
	router.Post("/reset-password", handlers.ResetPassword)

	return router
}
