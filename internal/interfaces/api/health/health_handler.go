package health

import (
	"net/http"

	"github.com/bubeha/PageInspectorBackend/pkg/response"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
}

func NewHealthHandler() *Handler {
	return &Handler{}
}

func (handler *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, "Alive!!!", http.StatusOK)
}

func (handler *Handler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handler.HealthCheck)

	return router
}
