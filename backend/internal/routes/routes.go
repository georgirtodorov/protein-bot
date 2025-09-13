// Package routes contains the HTTP routes for Protein Bot.
package routes

import (
	"database/sql"
	"net/http"

	"github.com/georgirtodorov/protein-bot/internal/handlers"
)

// Register wires URLs to handlers. Action-oriented routing design.
func Register(db *sql.DB) {
	http.HandleFunc("/", handlers.Welcome)

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		handlers.Add(w, r, db)
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		handlers.Status(w, r, db)
	})

	http.HandleFunc("/history", func(w http.ResponseWriter, r *http.Request) {
		handlers.History(w, r, db)
	})

	http.HandleFunc("/get-goal", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetGoal(w, r, db)
	})

	http.HandleFunc("/set-goal", func(w http.ResponseWriter, r *http.Request) {
		handlers.SetGoal(w, r, db)
	})

}
