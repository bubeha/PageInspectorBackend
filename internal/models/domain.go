package models

import (
	"time"

	"github.com/bubeha/PageInspectorBackend/internal/types"
	"github.com/google/uuid"
)

type Domain struct {
	ID        uuid.UUID          `db:"id" json:"id"`
	Name      string             `db:"name" json:"name" validate:"required,min=2,max=100"`
	BaseUrl   string             `db:"base_url" json:"base_url" validate:"required,domain"`
	Status    types.DomainStatus `db:"status" json:"status" validate:"required"`
	CreatedAt time.Time          `db:"created_at" json:"created_at"`
	UpdatedAt time.Time          `db:"updated_at" json:"updated_at"`
}
