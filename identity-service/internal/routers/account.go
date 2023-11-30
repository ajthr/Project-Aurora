package routers

import (
	"identity-service/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func managementRouter() *chi.Mux {
	router := chi.NewRouter()

	// route handlers
	router.Post("/change-mail", handlers.ChangeMail)
	router.Post("/refresh-token", handlers.RefreshToken)

	return router
}
