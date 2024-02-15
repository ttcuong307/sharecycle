package usecase

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"sharecycle/foundation/cerror"
	"sharecycle/internal/models"
	repository "sharecycle/internal/repository/db"
)

type UserInputPort interface {
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}

type User struct {
	ur repository.UserRepository
}

func NewUser(ur repository.UserRepository) UserInputPort {
	return &User{
		ur: ur,
	}
}

func (uc *User) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	// TODO: Check authentication before getting data. After implement Authentication and Authorization

	userInfo, err := uc.ur.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.User{}, cerror.WithCodeAndMessage(err, http.StatusNotFound, "userID not exist")
		}
		return &models.User{}, cerror.WithCodeAndMessage(err, http.StatusInternalServerError, "Cannot get User Info")
	}

	return userInfo, nil
}
