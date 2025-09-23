package main

import (
	"log"
	"os"

	"github.com/georgirtodorov/protein-bot/backend/internal/db"
	"github.com/georgirtodorov/protein-bot/backend/internal/routes"
	"github.com/georgirtodorov/protein-bot/backend/internal/server"
)

func main() {

	// Connect to the database
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	db, err := db.Connect(host, user, password, name, sslMode)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close() // close DB when main exits

	// Get the router (with routes + CORS middleware applied + DB passed)
	handler := routes.Register(db)

	// Start the server
	port := os.Getenv("PORT")
	server.Serve(port, handler)
}
