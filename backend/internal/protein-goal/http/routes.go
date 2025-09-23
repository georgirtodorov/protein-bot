// Package proteingoal provides HTTP routing for protein goal related endpoints.
package proteingoal

import (
	"database/sql"

	v1 "github.com/georgirtodorov/protein-bot/internal/protein/http/v1"
	"github.com/gorilla/mux"
)

// Register wires URLs to handlers. Action-oriented routing design.
func Register(r *mux.Router, db *sql.DB) {

	v1.Register(r, db)
}
