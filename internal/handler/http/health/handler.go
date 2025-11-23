package health

import "github.com/go-chi/chi/v5"

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (hh *HealthHandler) RegisterRoutes(router *chi.Mux) {
	router.Get("/health", HandlerFunc)
}
