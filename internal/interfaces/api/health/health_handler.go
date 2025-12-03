package health

import (
	"net/http"

	"github.com/bubeha/PageInspectorBackend/pkg/httputil"
	"github.com/bubeha/PageInspectorBackend/pkg/log"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	responder httputil.Responder
}

func NewHealthHandler() *Handler {
	return &Handler{
		responder: &httputil.JSONResponder{},
	}
}

func (handler *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	err := handler.responder.JSON(w, "Alive!!!", http.StatusOK)

	if err != nil {
		log.Infof("Failed to respond to healthcheck: %v", err)
	}
}

func (handler *Handler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/", handler.HealthCheck)

	return router
}
