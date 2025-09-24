package db

import (
	"database/sql"
	"time"
)

type ProteinGoal struct {
	ID        int        `json:"id"`
	Amount    int        `json:"amount"`
	CreatedAt time.Time  `json:"created_at"`
	Ended     *time.Time `json:"ended,omitempty"` // pointer so it can be null for last record
}

func SetProteinGoal(db *sql.DB, amount int) error {
	_, err := db.Exec("INSERT INTO protein_goals (amount) VALUES ($1)", amount)
	return err
}

func GetProteinGoal(db *sql.DB) (*ProteinGoal, error) {
	var goal ProteinGoal

	// Query the latest goal (DESC + LIMIT 1)
	err := db.QueryRow(`
        SELECT id, amount, created_at
        FROM protein_goals
        ORDER BY created_at DESC
        LIMIT 1
    `).Scan(&goal.ID, &goal.Amount, &goal.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, nil // No goal set
	}
	if err != nil {
		return nil, err
	}

	return &goal, nil
}

func GetProteinGoalHistory(db *sql.DB) ([]ProteinGoal, error) {
	// Query with DESC so latest goal comes first
	rows, err := db.Query("SELECT id, amount, created_at FROM protein_goals ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Close when done to avoid leaking resources

	var goals []ProteinGoal
	for rows.Next() {
		var g ProteinGoal
		if err := rows.Scan(&g.ID, &g.Amount, &g.CreatedAt); err != nil {
			return nil, err
		}
		goals = append(goals, g)
	}

	// Since we are in DESC order, "Ended" should point to the PREVIOUS row (older goal)
	for i := 0; i < len(goals)-1; i++ {
		goals[i].Ended = &goals[i+1].CreatedAt
	}

	return goals, nil
}
