package main

import (
	"context"
	"fmt"
	"sharecycle/configs"
	"sharecycle/internal/app"

	"github.com/caarlos0/env/v10"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Populate ENV
	err := configs.PopulateENV()
	if err != nil {
		fmt.Errorf("Parsing configuration err: %w", err)
	}

	// Config
	cfg := configs.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Errorf("Parsing configuration err: %w", err)
	}

	// Run
	app.Run(&cfg)
}
