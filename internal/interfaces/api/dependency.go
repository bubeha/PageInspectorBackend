package api

import (
	"github.com/bubeha/PageInspectorBackend/internal/app/domain"
	"github.com/bubeha/PageInspectorBackend/internal/config"
	"github.com/bubeha/PageInspectorBackend/internal/database"
	"github.com/bubeha/PageInspectorBackend/internal/repository"
	"github.com/bubeha/PageInspectorBackend/pkg/httputil"
)

type DataLayer struct {
	DomainRepo repository.DomainRepository
}

type Services struct {
	DomainService *domain.Service
}

type Infrastructure struct {
	DB        *database.DB
	Config    *config.Config
	Responser httputil.Responder
}
