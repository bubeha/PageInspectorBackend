package repository

import (
	"database/sql"
	"errors"

	"github.com/bubeha/PageInspectorBackend/internal/infrastructure/database"
	"github.com/bubeha/PageInspectorBackend/internal/infrastructure/log"
	"github.com/bubeha/PageInspectorBackend/internal/models"
	"github.com/google/uuid"
)

type PostgresDomainRepository struct {
	db *database.DB
}

type DomainRepository interface {
	FindAll() ([]models.Domain, error)
	FindByID(id uuid.UUID) (*models.Domain, error)
	Create(domain *models.Domain) error
}

func NewDomainRepository(db *database.DB) DomainRepository {
	return &PostgresDomainRepository{db: db}
}

func (r *PostgresDomainRepository) FindAll() ([]models.Domain, error) {
	var domains []models.Domain
	query := "SELECT * FROM domains ORDER BY created_at ASC"

	err := r.db.Select(&domains, query)

	if err != nil {
		return nil, err
	}

	return domains, nil
}

func (r *PostgresDomainRepository) FindByID(id uuid.UUID) (*models.Domain, error) {
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

func (r *PostgresDomainRepository) Create(domain *models.Domain) error {
	query := `INSERT INTO domains (name, base_url, status) VALUES (:name, :base_url, :status) RETURNING id`

	rows, err := r.db.NamedQuery(query, domain)

	if err != nil {
		log.Errorf("failed to insert domain: %s", err)

		return err
	}

	defer rows.Close()

	if rows.Next() {
		if rErr := rows.Scan(&domain.ID); rErr != nil {
			log.Error("failed to scan returned ID", "error", rErr)

			return rErr
		}
	}

	return nil
}
