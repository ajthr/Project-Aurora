package routers

import (
	"identity-service/internal/config"
	"identity-service/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func AccountManagementRouter(dbConfig *config.DBConfig) *chi.Mux {

	router := chi.NewRouter()

	// account handler
	accountHandler := handlers.NewAccountHandler(dbConfig.DB)

	router.Post("/change-mail", accountHandler.ChangeMail)
	router.Post("/refresh-token", accountHandler.RefreshToken)

	return router
}
