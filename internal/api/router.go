package api

import (
	"time"

	"github.com/bubeha/PageInspectorBackend/internal/config"
	"github.com/bubeha/PageInspectorBackend/internal/handler/http/domain"
	"github.com/bubeha/PageInspectorBackend/internal/handler/http/health"
	"github.com/bubeha/PageInspectorBackend/internal/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Setup(config *config.ServerConfig, repo *repository.DomainRepository) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Duration(config.Timeout) * time.Second))

	health.NewHealthHandler().RegisterRoutes(router)
	domain.NewDomainHandler(repo).RegisterRoutes(router)

	return router
}
