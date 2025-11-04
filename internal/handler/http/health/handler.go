package health

import "github.com/go-chi/chi/v5"

func RegisterRoutes(router *chi.Mux) {
	router.Get("/health", HandlerFunc)
}
