package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sharecycle/configs"
	"sharecycle/foundation/database"
	v1 "sharecycle/internal/controller/http/v1"
	"sharecycle/migrations"
	"sharecycle/pkg/logger"
	"syscall"
)

type Server struct {
	conf *configs.Config
	db   *database.DBV1
	l    logger.Logger
	h    *v1.Handler
}

func NewServer(conf *configs.Config, dbV1 *database.DBV1, l logger.Logger, h *v1.Handler) *Server {
	return &Server{
		conf: conf,
		db:   dbV1,
		l:    l,
		h:    h,
	}
}

func (r *Server) Run() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Migrates
	if err := migrations.Migrate(ctx, migrations.Config{
		User:           r.conf.DBConfigs.UserName,
		Password:       r.conf.DBConfigs.Password,
		Name:           r.conf.DBNames.V1,
		LoggerOverride: migrations.WrapLogger(r.l),
		DryRun:         false,
	}); err != nil {
		r.l.Info("migration errord")
	}

	// Init server
	sv := Init(Deps{
		DB:      r.db.DB,
		Logger:  r.l,
		APIAddr: r.conf.APIs.Address,
		Handler: r.h,
	})

	srv := &http.Server{
		Addr:    r.conf.APIs.Address,
		Handler: sv,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server
	go func() {
		r.l.Infof("Starting server on: %s", r.conf.APIs.Address)
		if err := srv.ListenAndServe(); err != nil {
			r.l.Info("server stopped")
		}
	}()

	// Graceful shutdown
	<-stop
	ctxShutdown, cancel := context.WithTimeout(context.Background(), r.conf.APIs.ShutdownTimeout)
	defer cancel()

	// Shutdown server
	if err := srv.Shutdown(ctxShutdown); err != nil {
		os.Exit(1)
	}

	r.l.Info("Server stopped gracefully")
}
