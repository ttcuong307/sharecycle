package main

import "sharecycle/internal/app"

func main() {
	// Run
	s := app.Ready()
	s.Run()
}
