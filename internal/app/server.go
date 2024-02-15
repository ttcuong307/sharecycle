package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sharecycle/configs"
	"sharecycle/foundation/database"
	"sharecycle/foundation/web"
	"sharecycle/migrations"
	"sharecycle/pkg/logger"
	"syscall"
)

type Server struct {
	conf *configs.Config
	db   *database.DBV1
}

func NewServer(conf *configs.Config, dbV1 *database.DBV1) *Server {
	return &Server{
		conf: conf,
		db:   dbV1,
	}
}

func (r *Server) Run() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Logger
	l, err := logger.NewArLogger(r.conf)
	if err != nil {
		fmt.Errorf("Server.Shutdown - Init sugar zap: %w", err)
	}
	l.Info("Init logger complete.")

	// Migrates
	if err = migrations.Migrate(ctx, migrations.Config{
		User:           r.conf.Database.UserName,
		Password:       r.conf.Database.Password,
		Name:           r.conf.Database.DBName,
		LoggerOverride: migrations.WrapLogger(l),
		DryRun:         false,
	}); err != nil {
		l.Info("migration errord")
	}

	// Init server
	sv := Init(web.Deps{
		DB:      r.db.DB,
		Logger:  l,
		APIAddr: r.conf.API.Address,
	})

	srv := &http.Server{
		Addr:    r.conf.API.Address,
		Handler: sv,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server
	go func() {
		l.Infof("Starting server on: %s", r.conf.API.Address)
		if err := srv.ListenAndServe(); err != nil {
			l.Info("server stopped")
		}
	}()

	// Graceful shutdown
	<-stop
	ctxShutdown, cancel := context.WithTimeout(context.Background(), r.conf.API.ShutdownTimeout)
	defer cancel()

	// Shutdown server
	if err := srv.Shutdown(ctxShutdown); err != nil {
		os.Exit(1)
	}

	l.Info("Server stopped gracefully")
}
