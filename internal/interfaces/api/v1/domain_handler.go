package v1

import (
	"net/http"

	"github.com/bubeha/PageInspectorBackend/internal/repository"
	"github.com/bubeha/PageInspectorBackend/pkg/httputil"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type DomainHandler struct {
	domainRepo repository.DomainRepository
	responser  httputil.Responder
}

func NewDomainHandler(domainRepo repository.DomainRepository, responder httputil.Responder) *DomainHandler {
	return &DomainHandler{
		domainRepo: domainRepo,
		responser:  responder,
	}
}

func (h *DomainHandler) Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{id}", h.ShowDomainHandlerFunc)

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

	domain, repoErr := h.domainRepo.FindByID(uid)

	if repoErr != nil {
		h.responser.Error(w, repoErr.Error(), http.StatusBadRequest)

		return
	}

	if err := h.responser.JSON(w, domain, http.StatusOK); err != nil {
		h.responser.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
