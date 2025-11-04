package api

import (
	"time"

	"github.com/bubeha/PageInspectorBackend/internal/config"
	"github.com/bubeha/PageInspectorBackend/internal/handler/http/health"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Setup(config *config.ServerConfig) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Duration(config.Timeout) * time.Second))

	health.RegisterRoutes(router)

	return router
}
