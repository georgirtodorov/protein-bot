// Package handlers contains the HTTP handlers for Protein Bot.
package handlers

import (
	"database/sql"
	"net/http"
)

// Register sets up all HTTP routes and injects the DB connection.
func Register(db *sql.DB) {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		add(w, r, db)
	})
	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		total(w, r, db)
	})
	http.HandleFunc("/history", func(w http.ResponseWriter, r *http.Request) {
		history(w, r, db)
	})
}

// welcome is a simple handler for the root endpoint.
func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Protein Bot!"))
}

// add handles the /add endpoint.
func add(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Example: use db to insert a protein record
	w.Write([]byte("Add endpoint"))
}

// total handles the /total endpoint.
func total(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Example: use db to calculate total protein intake
	w.Write([]byte("Total endpoint"))
}

// history handles the /history endpoint.
func history(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Example: use db to return protein intake history
	w.Write([]byte("History endpoint"))
}
