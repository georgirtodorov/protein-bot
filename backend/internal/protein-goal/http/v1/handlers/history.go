package goal

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/georgirtodorov/protein-bot/internal/db"
)

func GoalHistory(w http.ResponseWriter, r *http.Request, d *sql.DB) {
	// use db to fetch or update protein goal
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	result, err := db.GetProteinGoalHistory(d)
	if err != nil {
		http.Error(w, "Failed to fetch protein goal history", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
