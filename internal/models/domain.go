package models

import (
	"time"

	"github.com/bubeha/PageInspectorBackend/internal/types"
	"github.com/google/uuid"
)

type Domain struct {
	ID        uuid.UUID          `db:"id" json:"id"`
	Name      string             `db:"name" json:"name"`
	BaseUrl   string             `db:"base_url" json:"base_url"`
	Status    types.DomainStatus `db:"status" json:"status"`
	CreatedAt time.Time          `db:"created_at" json:"created_at"`
	UpdatedAt time.Time          `db:"updated_at" json:"updated_at"`
}
