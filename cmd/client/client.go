package main

import (
	"log"

	"github.com/dimk00z/GophKeeper/internal/app/client"
)

func main() {
	log.Println("Client CMD")
	client.StartClient()
}
