package main

import (
	"log"

	"github.com/tahahamdii/basic-api/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
