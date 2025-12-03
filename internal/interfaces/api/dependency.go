package api

import (
	"github.com/bubeha/PageInspectorBackend/internal/app/domain"
	"github.com/bubeha/PageInspectorBackend/internal/infrastructure/config"
	"github.com/bubeha/PageInspectorBackend/internal/infrastructure/database"
	"github.com/bubeha/PageInspectorBackend/internal/repository"
)

type DataLayer struct {
	DomainRepo repository.DomainRepository
}

type Services struct {
	DomainService *domain.Service
}

type Infrastructure struct {
	DB     *database.DB
	Config *config.Config
}
