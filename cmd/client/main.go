package main

import (
	"github.com/dimk00z/GophKeeper/internal/client/app"
	"github.com/dimk00z/GophKeeper/internal/client/app/build"
)

func main() {
	build.CheckBuild()
	// Configuration

	app.Execute()
}
