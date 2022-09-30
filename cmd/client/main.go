package main

import (
	"log"

	config "github.com/dimk00z/GophKeeper/config/client"
	"github.com/dimk00z/GophKeeper/internal/client/app"
	"github.com/dimk00z/GophKeeper/internal/client/app/build"
)

func main() {
	build.CheckBuild()
	// Configuration

	log.Println(config.LoadConfig())

	app.Execute()
}
