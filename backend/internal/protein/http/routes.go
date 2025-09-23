// Package protein provides HTTP routing for protein goal related endpoints.
package protein

import (
	"database/sql"

	v1 "github.com/georgirtodorov/protein-bot/internal/protein/http/v1"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router, db *sql.DB) {
	v1.Register(r, db) // now matches the signature
}
