// Package protein provides HTTP routing for protein goal related endpoints.
package protein

import (
	"database/sql"

	v1 "github.com/georgirtodorov/protein-bot/internal/protein/http/v1"
)

// Register wires URLs to handlers. Action-oriented routing design.
func Register(db *sql.DB) {

	v1.Register(db)
}
