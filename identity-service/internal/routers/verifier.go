package routers

import (
	"identity-service/internal/config"
	"identity-service/internal/middlewares"

	"github.com/go-chi/chi/v5"
)

func IdentityVerficationRouter() *chi.Mux {

	JWTConfig := config.NewJWTConfig()
	router := chi.NewRouter()

	router.Use(middlewares.Authenticator(*JWTConfig))

	// route handlers
	router.HandleFunc("/", nil)

	return router
}
