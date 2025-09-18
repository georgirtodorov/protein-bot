// Package v1 initialize and start the server
package v1

import (
	"database/sql"
	"net/http"
)

// Register wires URLs to handlers. Action-oriented routing design.
func Register(db *sql.DB) {

	http.HandleFunc("/v1/add", func(w http.ResponseWriter, r *http.Request) {
		Add(w, r, db)
	})

	http.HandleFunc("/v1/status", func(w http.ResponseWriter, r *http.Request) {
		Status(w, r, db)
	})

	http.HandleFunc("/v1/history", func(w http.ResponseWriter, r *http.Request) {
		History(w, r, db)
	})

}
