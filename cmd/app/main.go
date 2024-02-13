package main

import (
	"fmt"
	"sharecycle/configs"
	"sharecycle/internal/app"
)

func main() {

	// Populate ENV
	err := configs.PopulateENV()
	if err != nil {
		fmt.Errorf("Parsing configuration err: %w", err)
	}

	// Config
	cfg := configs.GetConfig()

	// Run
	app.Run(cfg)
}
