package database

import (
	"database/sql"
	"fmt"
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
		review_stage INTEGER DEFAULT 0,
		correct_answers INTEGER DEFAULT 0,
		incorrect_answers INTEGER DEFAULT 0
	);`

	_, err := db.Exec(flashcardsTable)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error while creating flashcards table")
	}
}
