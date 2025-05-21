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
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)`

	_, err := db.Exec(flashcardsTable)
	if err != nil {
		log.Fatal("Error while creating flashcards table")
	}
}
