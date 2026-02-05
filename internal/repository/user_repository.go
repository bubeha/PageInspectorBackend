package repository

import (
	"context"

	"github.com/bubeha/PageInspectorBackend/internal/models"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
}
