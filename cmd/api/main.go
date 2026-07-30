package main

import "log"

func main() {
	server, err := InitializeServer()
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}
