package v1

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/georgirtodorov/protein-bot/internal/db"
)

// Add handles the /add POST endpoint.
func Add(w http.ResponseWriter, r *http.Request, d *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	amountStr := r.FormValue("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		http.Error(w, "Invalid amount, number expected", http.StatusBadRequest)
		return
	}

	// Call the db layer to insert the protein record
	if err := db.InsertProtein(d, amount); err != nil {
		http.Error(w, "Failed to add entry", http.StatusInternalServerError)
		return
	}

	total, err := db.GetTotalForToday(d)
	if err != nil {
		http.Error(w, "Failed to fetch total", http.StatusInternalServerError)
		return
	}

	goal, err := db.GetProteinGoal(d)
	if err != nil {
		http.Error(w, "Failed to fetch goal", http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("Protein added: %d gr\nTotal: %d\nGoal: %d\nRemaining: %d",
		amount, total, goal, goal-total)
	w.Write([]byte(msg))
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

	fmt.Fprintf(w, "Total protein today: %d grams \n Goal: %d \n Remaining %d", total, goal, goal-total)
}

// History handles the /history endpoint.
func History(w http.ResponseWriter, r *http.Request, d *sql.DB) {
	entries, err := db.GetProteinHistory(d)
	if err != nil {
		http.Error(w, "Failed to fetch history", http.StatusInternalServerError)
		return
	}

	for _, e := range entries {
		fmt.Fprintf(w, "%s: %d grams\n", e.CreatedAt.Format("2006-01-02 15:04"), e.Amount)
	}
}
