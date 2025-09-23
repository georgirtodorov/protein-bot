// Package routes contains the HTTP routes for Protein Bot.
package routes

import (
	"database/sql"
	"net/http"
	"os"

	proteingoal "github.com/georgirtodorov/protein-bot/internal/protein-goal/http"
	protein "github.com/georgirtodorov/protein-bot/internal/protein/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	go_handlers "github.com/georgirtodorov/protein-bot/internal/handlers"
)

// Register wires URLs to handlers. Action-oriented routing design.
func Register(db *sql.DB) http.Handler {
	r := mux.NewRouter() // top-level router

	apiOrigin := os.Getenv("REACT_APP_API_URL") // <- dynamic origin

	h := handlers.CORS(
		handlers.AllowedOrigins([]string{apiOrigin}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	proteingoal.Register(r, db)
	protein.Register(r, db) // protein v1 routes

	// Optional welcome route
	r.HandleFunc("/", go_handlers.Welcome)

	return h(r) // return handler for ListenAndServe
}
