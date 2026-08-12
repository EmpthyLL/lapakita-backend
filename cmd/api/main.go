package main

import (
	"lapakita-backend/pkg/i18n"
	"log"
)

func main() {
	i18n.Init()
	server, err := InitializeServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
