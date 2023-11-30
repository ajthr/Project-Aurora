package routers

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.StripSlashes)
	router.Use(httprate.LimitByIP(100, 1*time.Minute))

	// route handlers
	router.HandleFunc("/", nil)

	return router
}
