package routers

import (
	"identity-service/internal/config"
	"identity-service/internal/middlewares"

	"github.com/go-chi/chi/v5"
)

func IdentityVerficationRouter(jwtConfig *config.JWTConfig) *chi.Mux {

	router := chi.NewRouter()

	router.Use(middlewares.Authenticator(jwtConfig))

	// route handlers
	router.HandleFunc("/", nil)

	return router
}
