package api

import (
	"net/http"
	"time"

	"github.com/bubeha/PageInspectorBackend/internal/interfaces/api/health"
	v1 "github.com/bubeha/PageInspectorBackend/internal/interfaces/api/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	port   string
	router *chi.Mux
}

func NewServer(data *DataLayer, services *Services, infra *Infrastructure) *Server {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(time.Duration(infra.Config.Server.Timeout) * time.Second))

	// Handlers
	router.Mount("/health", health.NewHealthHandler().Routes())
	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/domains", v1.NewDomainHandler(data.DomainRepo, infra.Responser, *services.DomainService).Routes())
	})

	return &Server{
		router: router,
		port:   infra.Config.Server.Port,
	}
}

func (s *Server) Run() error {
	return http.ListenAndServe(":"+s.port, s.router)
}
