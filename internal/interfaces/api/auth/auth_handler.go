package auth

import (
	"net/http"
	"time"

	"github.com/bubeha/PageInspectorBackend/internal/app/auth"
	"github.com/bubeha/PageInspectorBackend/internal/infrastructure/config"
	"github.com/bubeha/PageInspectorBackend/pkg/log"
	"github.com/bubeha/PageInspectorBackend/pkg/response"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	config  *config.JWTConfig
	service *auth.LoginService
}

func NewAuthHandler(cnf *config.JWTConfig, service *auth.LoginService) *Handler {
	return &Handler{
		config:  cnf,
		service: service,
	}
}

func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	pairs, err := handler.service.Login(r.Context(), "example.com", "secret")

	if err != nil {
		log.Errorf("Failed to generate tokens: %v", err)
		response.Error(w, "Server error", 500)

		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "refresh token value",
		Expires:  time.Now().Add(handler.config.PublicTTL),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	response.JSON(w, struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}{
		AccessToken: pairs.AccessToken,
		ExpiresIn:   900,
	}, 200)
}

func (handler *Handler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/login", handler.Login)

	return router
}
