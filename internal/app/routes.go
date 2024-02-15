package app

import (
	"gorm.io/gorm"
	"net/http"
	"sharecycle/foundation/web"
	v1 "sharecycle/internal/controller/http/v1"
	"sharecycle/pkg/logger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type Deps struct {
	DB      *gorm.DB
	Logger  logger.Logger
	APIAddr string
	Handler *v1.Handler
}

type Router struct {
	Mux chi.Router
}

func (r *Router) Use(middleware ...func(handler http.Handler) http.Handler) {
	r.Mux.Use(middleware...)
}

func Init(d Deps) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set CORS
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	r.Use(corsOptions.Handler)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) { web.Respond(w, "OK", http.StatusOK) })
		v1.RegisterRoutes(r, *d.Handler)
	})

	return r
}
