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

	r.HandleFunc("/proteingoal/set", func(w http.ResponseWriter, r *http.Request) {
		goal.SetGoal(w, r, db)
	}).Methods("POST")
	r.HandleFunc("/proteingoal/get", func(w http.ResponseWriter, r *http.Request) {
		goal.GetGoal(w, r, db)
	}).Methods("GET")
}
