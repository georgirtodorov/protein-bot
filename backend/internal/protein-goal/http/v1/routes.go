// Package v1 initialize and start the server
package v1

import (
	"database/sql"
	"net/http"
)

// Register wires URLs to handlers. Action-oriented routing design.
func Register(db *sql.DB) {

	http.HandleFunc("/v1/get-goal", func(w http.ResponseWriter, r *http.Request) {
		GetGoal(w, r, db)
	})

	http.HandleFunc("/v1/set-goal", func(w http.ResponseWriter, r *http.Request) {
		SetGoal(w, r, db)
	})

}
