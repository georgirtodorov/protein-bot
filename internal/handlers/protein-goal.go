package handlers

import (
	"database/sql"
	"net/http"
)

func Goal(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// use db to fetch or update protein goal
	w.Write([]byte("Goal endpoint"))
}
