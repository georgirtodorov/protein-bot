package protein

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/georgirtodorov/protein-bot/internal/db"
)

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
