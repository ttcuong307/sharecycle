package web

import (
	"net/http"
	"sharecycle/pkg/logger"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

type Deps struct {
	DB      *gorm.DB
	Logger  logger.Logger
	APIAddr string
}

type Router struct {
	Mux chi.Router
}

func (r *Router) Use(middleware ...func(handler http.Handler) http.Handler) {
	r.Mux.Use(middleware...)
}
