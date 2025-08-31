package db

import (
	"database/sql"
	"time"
)

type ProteinEntry struct {
	ID        int       `json:"id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func InsertProtein(db *sql.DB, amount int) error {
	_, err := db.Exec(
		"INSERT INTO protein_entries (amount) VALUES ($1)",
		amount,
	)
	return err
}

func GetTotalForToday(db *sql.DB) (int, error) {
	var total int
	err := db.QueryRow(`
		SELECT COALESCE(SUM(amount), 0)
		FROM protein_entries
		WHERE DATE(created_at) = CURRENT_DATE
	`).Scan(&total)
	return total, err
}

func GetProteinHistory(d *sql.DB) ([]ProteinEntry, error) {
	rows, err := d.Query("SELECT id, amount, created_at FROM protein_entry ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []ProteinEntry
	for rows.Next() {
		var e ProteinEntry
		if err := rows.Scan(&e.ID, &e.Amount, &e.CreatedAt); err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}
