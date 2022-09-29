package main

import (
	"log"

	client "github.com/dimk00z/GophKeeper/internal/client/app"
)

func main() {
	log.Println("Client CMD")
	client.StartClient()
}
