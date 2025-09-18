// Package routes contains the HTTP routes for Protein Bot.
package routes

import (
	"database/sql"
	"net/http"

	proteingoal "github.com/georgirtodorov/protein-bot/internal/protein-goal/http"
	protein "github.com/georgirtodorov/protein-bot/internal/protein/http"

	"github.com/georgirtodorov/protein-bot/internal/handlers"
)

// Register wires URLs to handlers. Action-oriented routing design.
func Register(db *sql.DB) {

	proteingoal.Register(db)

	protein.Register(db)

	http.HandleFunc("/", handlers.Welcome)
}
