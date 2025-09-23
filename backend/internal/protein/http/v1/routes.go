// Package v1 initialize and start the server
package v1

import (
	"database/sql"
	"net/http"

	protein "github.com/georgirtodorov/protein-bot/internal/protein/http/v1/handlers"

	"github.com/gorilla/mux"
)

// Register wires URLs to handlers. Action-oriented routing design.
func Register(r *mux.Router, db *sql.DB) {

	r.HandleFunc("/v1/add", func(w http.ResponseWriter, r *http.Request) {
		protein.Add(w, r, db)
	}).Methods("POST")

	r.HandleFunc("/v1/status", func(w http.ResponseWriter, r *http.Request) {
		protein.Status(w, r, db)
	}).Methods("GET")

	r.HandleFunc("/v1/history", func(w http.ResponseWriter, r *http.Request) {
		protein.History(w, r, db)
	}).Methods("GET")
}
