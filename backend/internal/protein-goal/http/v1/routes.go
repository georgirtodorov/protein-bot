// Package v1 initialize and start the server
package v1

import (
	"database/sql"
	"net/http"

	goal "github.com/georgirtodorov/protein-bot/internal/protein-goal/http/v1/handlers"

	"github.com/gorilla/mux"
)

// Register wires URLs to handlers. Action-oriented routing design.
func Register(r *mux.Router, db *sql.DB) {

	r.HandleFunc("/v1/protein/goal", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			goal.GetGoal(w, r, db)
		case http.MethodPost:
			goal.SetGoal(w, r, db)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	r.HandleFunc("/v1/protein/goal/history", func(w http.ResponseWriter, r *http.Request) {
		goal.GoalHistory(w, r, db)
	}).Methods("GET")
}
