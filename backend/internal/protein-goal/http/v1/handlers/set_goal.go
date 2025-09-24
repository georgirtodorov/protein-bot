package goal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/georgirtodorov/protein-bot/internal/db"
)

func SetGoal(w http.ResponseWriter, r *http.Request, d *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON body
	var payload struct {
		Amount int `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Set the goal in DB
	if err := db.SetProteinGoal(d, payload.Amount); err != nil {
		http.Error(w, "Failed to set goal", http.StatusInternalServerError)
		return
	}

	fmt.Println("Protein goal set:", payload.Amount, "grams daily")

	// Respond
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}
