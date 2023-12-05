package routers

import (
	"time"

	"identity-service/internal/config"
	"identity-service/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func NewRouter(dbConfig *config.DBConfig, jwtConfig *config.JWTConfig, googleClient *config.GoogleAuthClient, mailClient *config.MailConfig) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.StripSlashes)
	router.Use(httprate.Limit(
		25,
		1*time.Second,
		httprate.WithLimitHandler(handlers.TooManyRequests)))

	router.NotFound(handlers.NotFound)
	router.MethodNotAllowed(handlers.MethodNotAllowed)

	// mount subroutes
	router.Mount("/auth", AuthRouter(dbConfig, googleClient, mailClient, jwtConfig))
	router.Mount("/user", AccountManagementRouter(dbConfig))
	router.Mount("/verify-token", IdentityVerficationRouter(jwtConfig))

	return router
}
