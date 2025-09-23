// Package protein initialize and start the server
package protein

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/georgirtodorov/protein-bot/internal/db"
)

type AddResponse struct {
	Added     int `json:"added"`
	Total     int `json:"total"`
	Goal      int `json:"goal"`
	Remaining int `json:"remaining"`
}

// Add handles the /add POST endpoint.
func Add(w http.ResponseWriter, r *http.Request, d *sql.DB) {
	fmt.Println("Add handler hit")

	if r.Method != http.MethodPost {
		fmt.Println("Invalid method", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type AddRequest struct {
		Amount int `json:"amount"`
	}

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Failed to decode JSON:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	amount := req.Amount
	fmt.Println("Received amount:", amount)

	if err := db.InsertProtein(d, amount); err != nil {
		fmt.Println("DB insert failed:", err)
		http.Error(w, "Failed to add entry", http.StatusInternalServerError)
		return
	}

	total, err := db.GetTotalForToday(d)
	if err != nil {
		fmt.Println("DB total failed:", err)
		http.Error(w, "Failed to fetch total", http.StatusInternalServerError)
		return
	}

	goal, err := db.GetProteinGoal(d)
	if err != nil {
		fmt.Println("DB goal failed:", err)
		http.Error(w, "Failed to fetch goal", http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("Protein added: %d gr\nTotal: %d\nGoal: %d\nRemaining: %d",
		amount, total, goal, goal-total)
	fmt.Println("Sending response:", msg)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(AddResponse{
		Added:     amount,
		Total:     total,
		Goal:      goal,
		Remaining: goal - total,
	})
}
