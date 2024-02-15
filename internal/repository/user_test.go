package repository_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"sharecycle/internal/models"
	mock_repository "sharecycle/internal/repository/mock"
	"testing"
)

func TestUserRepository_GetUserByID(t *testing.T) {
	ctx := context.Background()

	t.Run("FindByID", func(t *testing.T) {
		mockCtl := gomock.NewController(t)
		defer mockCtl.Finish()

		repo := mock_repository.NewMockUserRepository(mockCtl)

		user := &models.User{
			ID:     "1",
			Name:   "test1",
			Gender: "1",
		}

		repo.EXPECT().GetUserByID(ctx, "1").Return(user, nil)

		user, err := repo.GetUserByID(ctx, "1")
		require.NoError(t, err)
		require.Equal(t, "1", user.ID)
	})
}
