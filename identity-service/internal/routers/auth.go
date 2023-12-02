package routers

import (
	"identity-service/internal/config"
	"identity-service/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(dbConfig *config.DBConfig) *chi.Mux {
	router := chi.NewRouter()

	// auth handler
	authHandler := handlers.NewAuthHandler(dbConfig.DB)

	router.Post("/signin", authHandler.SignIn)
	router.Post("/signup", authHandler.SignUp)
	router.Post("/google-signin", authHandler.GoogleSignIn)
	router.Post("/reset-password", authHandler.ResetPassword)

	return router
}
