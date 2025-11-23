package domain

import "github.com/bubeha/PageInspectorBackend/internal/repository"

type Service struct {
	repo repository.DomainRepository
}

func NewDomainService(repo repository.DomainRepository) *Service {
	return &Service{repo: repo}
}
