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

func Run(cfg *configs.Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Logger
	l, err := logger.NewArLogger(cfg)
	if err != nil {
		fmt.Errorf("Server.Shutdown - Init sugar zap: %w", err)
	}
	l.Info("Init logger complete.")

	// Migrates
	if err = migrations.Migrate(ctx, migrations.Config{
		User:           cfg.Database.UserName,
		Password:       cfg.Database.Password,
		Name:           cfg.Database.DBName,
		LoggerOverride: migrations.WrapLogger(l),
		DryRun:         false,
	}); err != nil {
		l.Info("migration errord")
	}

	// Start Database
	db := database.NewDB(cfg)

	// Init server
	r := Init(web.Deps{
		DB:      db.DB,
		Logger:  l,
		APIAddr: cfg.API.Address,
	})

	srv := &http.Server{
		Addr:    cfg.API.Address,
		Handler: r,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server
	go func() {
		l.Infof("Starting server on: %s", cfg.API.Address)
		if err := srv.ListenAndServe(); err != nil {
			l.Info("server stopped")
		}
	}()

	// Graceful shutdown
	<-stop
	ctxShutdown, cancel := context.WithTimeout(context.Background(), cfg.API.ShutdownTimeout)
	defer cancel()

	// Shutdown server
	if err := srv.Shutdown(ctxShutdown); err != nil {
		os.Exit(1)
	}

	l.Info("Server stopped gracefully")
}
