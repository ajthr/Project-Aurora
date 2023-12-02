package routers

import (
	"log"
	"os"
	"time"

	"identity-service/internal/config"
	"identity-service/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

var (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_NAME")
)

func NewRouter() *chi.Mux {

	dbConfig, err := config.NewDBConfig(host, port, user, password, dbname)

	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	router := chi.NewRouter()

	router.Use(middleware.StripSlashes)
	router.Use(httprate.Limit(
		25,
		1*time.Second,
		httprate.WithLimitHandler(handlers.TooManyRequests)))

	router.NotFound(handlers.NotFound)
	router.MethodNotAllowed(handlers.MethodNotAllowed)

	// mount subroutes
	router.Mount("/auth", AuthRouter(dbConfig))
	router.Mount("/user", AccountManagementRouter(dbConfig))
	router.Mount("/verify-token", IdentityVerficationRouter())

	return router
}
