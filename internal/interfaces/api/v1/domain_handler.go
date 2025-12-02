package v1

import (
	"encoding/json"
	"net/http"

	"github.com/bubeha/PageInspectorBackend/internal/app/domain"
	"github.com/bubeha/PageInspectorBackend/internal/models"
	"github.com/bubeha/PageInspectorBackend/internal/repository"
	"github.com/bubeha/PageInspectorBackend/pkg/httputil"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type DomainHandler struct {
	domainRepo    repository.DomainRepository
	responser     httputil.Responder
	domainService domain.Service
}

func NewDomainHandler(domainRepo repository.DomainRepository, responder httputil.Responder, service domain.Service) *DomainHandler {
	return &DomainHandler{
		domainRepo:    domainRepo,
		responser:     responder,
		domainService: service,
	}
}

func (h *DomainHandler) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", h.ShowDomainHandlerFunc)
	router.Post("/", h.CreateDomainHandlerFunc)

	return router
}

func (h *DomainHandler) ShowDomainHandlerFunc(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		h.responser.Error(w, "Id is required", http.StatusBadRequest)
	}

	uid, uuidErr := uuid.Parse(id)

	if uuidErr != nil {
		h.responser.Error(w, uuidErr.Error(), http.StatusBadRequest)

		return
	}

	entry, repoErr := h.domainRepo.FindByID(uid)

	if repoErr != nil {
		h.responser.Error(w, repoErr.Error(), http.StatusBadRequest)

		return
	}

	if err := h.responser.JSON(w, entry, http.StatusOK); err != nil {
		h.responser.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *DomainHandler) CreateDomainHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var entry models.Domain

	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		h.responser.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	if err := h.domainService.CreateDomain(&entry); err != nil {
		h.responser.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := h.responser.JSON(w, entry, http.StatusCreated); err != nil {
		h.responser.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
