package protein

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/georgirtodorov/protein-bot/internal/db"
)

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
