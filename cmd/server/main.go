package main

import (
	"log"

	config "github.com/dimk00z/GophKeeper/config/server"
	server "github.com/dimk00z/GophKeeper/internal/app/server"
)

func main() {
	log.Println("Server App")

	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	server.Run(cfg)
}
