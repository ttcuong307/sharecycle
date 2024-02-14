package main

import (
	"sharecycle/configs"
	"sharecycle/internal/app"
)

func main() {
	// Config
	cfg := configs.GetConfig()

	// Run
	app.Run(cfg)
}
