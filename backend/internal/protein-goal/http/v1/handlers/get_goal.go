// Package goal get and set
package goal

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/georgirtodorov/protein-bot/internal/db"
)

func GetGoal(w http.ResponseWriter, r *http.Request, d *sql.DB) {
	// use db to fetch or update protein goal
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	total, err := db.GetProteinGoal(d)
	if err != nil {
		http.Error(w, "Failed to fetch protein goal", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Total protein today: %d grams", total)
}
