-- +migrate Up
CREATE TABLE IF NOT EXISTS protein_entries (
    id SERIAL PRIMARY KEY,
    amount INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
