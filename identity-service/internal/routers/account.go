package routers

import (
	"identity-service/internal/config"
	"identity-service/internal/handlers"
	"identity-service/internal/middlewares"

	"github.com/go-chi/chi/v5"
)

func managementRouter() *chi.Mux {

	JWTConfig := config.NewJWTConfig()
	router := chi.NewRouter()

	router.Use(middlewares.Authenticator(*JWTConfig))

	// route handlers
	router.Post("/change-mail", handlers.ChangeMail)
	router.Post("/refresh-token", handlers.RefreshToken)

	return router
}
