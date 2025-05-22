package database

import (
	"database/sql"
	"log"
)

func CreateTables(db *sql.DB) {
	flashcardsTable := `
		CREATE TABLE IF NOT EXISTS flashcards (
		id SERIAL PRIMARY KEY,
		front VARCHAR(255) NOT NULL,
		back VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		last_review TIMESTAMP NULL,
		correct_answers INTEGER DEFAULT 0,
		tries INTEGER DEFAULT 0
	);`

	_, err := db.Exec(flashcardsTable)
	if err != nil {
		log.Fatal("Error while creating flashcards table")
	}
}
