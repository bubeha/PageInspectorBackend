package repository

import (
	"database/sql"
	"errors"

	"github.com/bubeha/PageInspectorBackend/internal/database"
	"github.com/bubeha/PageInspectorBackend/internal/models"
	"github.com/google/uuid"
)

type DomainRepository struct {
	db *database.DB
}

func NewDomainRepository(db *database.DB) *DomainRepository {
	return &DomainRepository{db: db}
}

func (r *DomainRepository) FindAll() ([]models.Domain, error) {
	var domains []models.Domain
	query := "SELECT * FROM domains ORDER BY created_at ASC"

	err := r.db.Select(&domains, query)

	if err != nil {
		return nil, err
	}

	return domains, nil
}

func (r *DomainRepository) FindByID(id uuid.UUID) (*models.Domain, error) {
	var domain models.Domain
	query := "SELECT * FROM domains WHERE id = $1 LIMIT 1;"

	err := r.db.Get(&domain, query, id)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("domain not found")
	}

	if err != nil {
		return nil, err
	}

	return &domain, nil
}
