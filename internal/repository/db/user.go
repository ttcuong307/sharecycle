package repository

import (
	"context"
	"sharecycle/internal/models"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}
