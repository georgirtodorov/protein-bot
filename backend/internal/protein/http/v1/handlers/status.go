package protein

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/georgirtodorov/protein-bot/internal/db"
)

type StatusResponse struct {
	Total     int `json:"total"`
	Goal      int `json:"goal"`
	Remaining int `json:"remaining"`
}

func Status(w http.ResponseWriter, r *http.Request, d *sql.DB) {
	total, err := db.GetTotalForToday(d)
	if err != nil {
		http.Error(w, "Failed to fetch total", http.StatusInternalServerError)
		return
	}

	goal, err := db.GetProteinGoal(d)
	if err != nil {
		http.Error(w, "Failed to fetch total", http.StatusInternalServerError)
		return
	}

	response := StatusResponse{
		Total:     total,
		Goal:      goal,
		Remaining: goal - total,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
