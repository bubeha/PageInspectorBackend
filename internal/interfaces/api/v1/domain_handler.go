package v1

import (
	"net/http"

	"github.com/bubeha/PageInspectorBackend/internal/app/domain"
	"github.com/bubeha/PageInspectorBackend/internal/models"
	"github.com/bubeha/PageInspectorBackend/internal/repository"
	"github.com/bubeha/PageInspectorBackend/pkg/request"
	"github.com/bubeha/PageInspectorBackend/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type DomainHandler struct {
	domainRepo    repository.DomainRepository
	domainService domain.Service
}

func NewDomainHandler(domainRepo repository.DomainRepository, service domain.Service) *DomainHandler {
	return &DomainHandler{
		domainRepo:    domainRepo,
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
		response.Error(w, "Id is required", http.StatusBadRequest)
	}

	uid, uuidErr := uuid.Parse(id)

	if uuidErr != nil {
		response.Error(w, uuidErr.Error(), http.StatusBadRequest)

		return
	}

	entry, repoErr := h.domainRepo.FindByID(uid)

	if repoErr != nil {
		response.Error(w, repoErr.Error(), http.StatusBadRequest)

		return
	}

	response.JSON(w, entry, http.StatusOK)
}

func (h *DomainHandler) CreateDomainHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var entry models.Domain

	if err := request.JSON(r, &entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := h.domainService.CreateDomain(&entry); err != nil {
		response.JsonError(w, err, http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response.JSON(w, entry, http.StatusCreated)
}
