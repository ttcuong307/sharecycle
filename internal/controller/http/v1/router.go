package v1

import (
	"github.com/go-chi/chi"
)

func RegisterRoutes(r chi.Router, h Handler) {
	r.Get("/users/{userId}", h.GetUserByID)
}
