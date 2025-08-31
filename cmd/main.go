package main

import (
	"log"
	"os"

	"github.com/georgirtodorov/protein-bot/internal/db"
	"github.com/georgirtodorov/protein-bot/internal/routes"
	"github.com/georgirtodorov/protein-bot/internal/server"
)

func main() {

	// Connect to the database
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	db, err := db.Connect(host, user, password, name)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close() // close DB when main exits

	// Register all routes with DB
	routes.Register(db)

	// Start the server
	port := os.Getenv("PORT")
	server.Serve(port)
}
