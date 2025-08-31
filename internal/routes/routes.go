// Package routes contains the HTTP routes for Protein Bot.
package routes

import (
	"database/sql"
	"net/http"

	"github.com/georgirtodorov/protein-bot/internal/handlers"
)

// Register wires URLs to handlers
func Register(db *sql.DB) {
	http.HandleFunc("/", handlers.Welcome)
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		handlers.Add(w, r, db)
	})
	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		handlers.Total(w, r, db)
	})
	http.HandleFunc("/history", func(w http.ResponseWriter, r *http.Request) {
		handlers.History(w, r, db)
	})
	http.HandleFunc("/goal", func(w http.ResponseWriter, r *http.Request) {
		handlers.Goal(w, r, db)
	})
}
