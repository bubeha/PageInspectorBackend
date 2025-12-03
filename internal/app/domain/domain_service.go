package domain

import (
	"fmt"

	"github.com/bubeha/PageInspectorBackend/internal/models"
	"github.com/bubeha/PageInspectorBackend/internal/repository"
	"github.com/bubeha/PageInspectorBackend/internal/types"
	"github.com/bubeha/PageInspectorBackend/pkg/validator"
)

type Service struct {
	repo repository.DomainRepository
}

func NewDomainService(repo repository.DomainRepository) *Service {
	return &Service{repo: repo}
}

func (service *Service) CreateDomain(domain *models.Domain) error {
	domain.Status = types.DomainStatusCreated

	if err := validator.Validate(domain); err != nil {
		return err
	}

	exists, err := service.repo.ExistByUrl(domain.BaseUrl)

	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("domain already exists")
	}

	return service.repo.Create(domain)
}
