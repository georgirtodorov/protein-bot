package main

import (
	"log"

	"github.com/georgirtodorov/protein-bot/internal/app"
	"github.com/georgirtodorov/protein-bot/internal/db"
	"github.com/georgirtodorov/protein-bot/internal/server"
)

func main() {
	d, err := db.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer d.Close() // close DB when main exits

	app := &app.App{DB: d} // pass DB to App

	app.Routes() // register routes

	server.Start() // start the server
}
