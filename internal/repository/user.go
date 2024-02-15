package repository

import (
	"context"
	"sharecycle/foundation/database"
	"sharecycle/internal/models"
)

type UserRepository struct {
	BaseRepository
}

func NewUserRepository(db *database.DBV1) *UserRepository {
	return &UserRepository{BaseRepository{DB: db.DB}}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User

	tx := r.WithTx(ctx).
		Where("id = ?", id)

	if err := tx.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
