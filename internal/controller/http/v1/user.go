package v1

import (
	"errors"
	"github.com/go-chi/chi"
	"net/http"
	"sharecycle/foundation/web"
)

type UserHandler interface {
	GetUserByID(w http.ResponseWriter, r *http.Request)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId := chi.URLParam(r, "userId")
	if userId == "" {
		h.Log.Error("http - v1 - GetUserByID - parse userId")
		web.RespondError(w, errors.New("parse userId"))
		return
	}

	userInfo, err := h.User.GetUserByID(ctx, userId)
	if err != nil {
		h.Log.Error(err, "http - v1 - GetUserByID - parse userId")
		web.RespondError(w, err)
		return
	}

	web.Respond(w, userInfo, http.StatusOK)
	return
}
