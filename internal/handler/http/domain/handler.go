package domain

import (
	"github.com/bubeha/PageInspectorBackend/internal/repository"
	"github.com/bubeha/PageInspectorBackend/pkg/httputil"
	"github.com/go-chi/chi/v5"
)

type DomainHandler struct {
	repo      *repository.DomainRepository
	responser httputil.Responder
}

func NewDomainHandler(repo *repository.DomainRepository) *DomainHandler {
	return &DomainHandler{repo: repo, responser: &httputil.JSONResponder{}}
}

func (dh *DomainHandler) RegisterRoutes(router *chi.Mux) {
	router.Get("/domains/{id}", dh.ShowDomainHandlerFunc)
}
