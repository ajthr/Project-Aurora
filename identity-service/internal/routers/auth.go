package routers

import (
	"identity-service/internal/config"
	"identity-service/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func AuthRouter(dbConfig *config.DBConfig, googleClient *config.GoogleAuthClient, mailConfig *config.MailConfig, jwtConfig *config.JWTConfig) *chi.Mux {
	router := chi.NewRouter()

	// auth handler
	authHandler := handlers.NewAuthHandler(dbConfig.DB, googleClient, mailConfig, jwtConfig)

	router.Post("/signin", authHandler.SignIn)
	router.Post("/validate-otp", authHandler.ValidateOtp)
	router.Post("/google-signin", authHandler.GoogleSignIn)

	return router
}
