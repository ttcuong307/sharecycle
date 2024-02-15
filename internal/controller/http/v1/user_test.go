package v1_test

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	v1 "sharecycle/internal/controller/http/v1"
	"sharecycle/internal/models"
	mock_usecase "sharecycle/internal/usecase/mock"
	"testing"
)

func TestHandler_GetUserByID(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// TODO: Check authentication
		mock_usecase := mock_usecase.NewMockUserInputPort(ctrl)

		userOutput := &models.User{
			ID:     "1",
			Name:   "test1",
			Gender: "1",
		}

		mock_usecase.EXPECT().GetUserByID(gomock.Any(), "1").Return(userOutput, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/users/{userId}", nil)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("userId", "1")

		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

		handler := &v1.Handler{
			User: mock_usecase,
		}

		handler.GetUserByID(w, r)
		require.Equal(t, http.StatusOK, w.Code)

		user := models.User{}
		err := json.Unmarshal(w.Body.Bytes(), &user)
		require.NoError(t, err)
		require.Equal(t, userOutput.ID, user.ID)
	})
}
