package main

import (
	"log"

	config "github.com/dimk00z/GophKeeper/config/server"
	server "github.com/dimk00z/GophKeeper/internal/app/server"
)

// @title Gophkeeper Server
// @version 1.0.0
// @description Gophkeeper project
// @contact.name Kuznetsov Dmitriy
// @contact.url https://github.com/dimk00z
// @contact.email dimk0z@yandex.ru
// @host localhost:8080
// @BasePath /api/v1
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
