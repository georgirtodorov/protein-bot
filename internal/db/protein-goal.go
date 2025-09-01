package db

import (
	"database/sql"
	"time"
)

type ProteinGoal struct {
	ID        int       `json:"id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func SetProteinGoal(db *sql.DB, amount int) error {
	_, err := db.Exec(
		"INSERT INTO protein_goals (amount) VALUES ($1) ON CONFLICT (id) DO UPDATE SET amount = EXCLUDED.amount",
		amount,
	)
	return err
}

func GetProteinGoal(db *sql.DB) (int, error) {
	var goal int
	err := db.QueryRow("SELECT amount FROM protein_goals ORDER BY created_at DESC LIMIT 1").Scan(&goal)
	if err == sql.ErrNoRows {
		return 0, nil // No goal set
	}
	return goal, err
}
