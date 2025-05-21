package flashcards

import "database/sql"

type FlashcardRepository struct {
	DB *sql.DB
}

func NewFlashcardRepository(db *sql.DB) *FlashcardRepository {
	return &FlashcardRepository{DB: db}
}

