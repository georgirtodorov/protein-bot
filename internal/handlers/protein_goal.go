// Package handlers contains the HTTP handlers for Protein Bot.

package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

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

func SetGoal(w http.ResponseWriter, r *http.Request, d *sql.DB) {
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
	if err := db.SetProteinGoal(d, amount); err != nil {
		http.Error(w, "Failed to add entry", http.StatusInternalServerError)
		return
	}

	msg := fmt.Sprintf("Protein goal set: %d grams daily", amount)
	w.Write([]byte(msg))
}
