package routers

import (
	"identity-service/internal/handlers"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.StripSlashes)
	router.Use(httprate.Limit(
		25,
		1*time.Second,
		httprate.WithLimitHandler(handlers.TooManyRequests)))

	router.NotFound(handlers.NotFound)
	router.MethodNotAllowed(handlers.MethodNotAllowed)

	// mount subroutes
	router.Mount("/", AuthRouter())
	router.Mount("/me", AccountManagementRouter())
	router.Mount("/verify-token", IdentityVerficationRouter())

	return router
}
