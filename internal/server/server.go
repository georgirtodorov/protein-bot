// Package server initialize and start the server
package server

import (
	"log"
	"net/http"
	"os"
)

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalf("server failed to start, PORT environment variable not set")
	}
	addr := ":" + port
	log.Printf("Starting running on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
