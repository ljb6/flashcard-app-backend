package flashcards

import (
	"database/sql"
	"errors"
)

type FlashcardRepository struct {
	DB *sql.DB
}

func NewFlashcardRepository(db *sql.DB) *FlashcardRepository {
	return &FlashcardRepository{DB: db}
}

func (r *FlashcardRepository) CreateFlashcard(front, back string) error {
	query := `INSERT INTO flashcards (front, back) VALUES ($1, $2)`
	_, err := r.DB.Exec(query, front, back)
	if err != nil {
		return err
	}
	return nil
}

func (r *FlashcardRepository) GetFlashcards() ([]Flashcard, error) {
	query := `SELECT id, front, back, created_at FROM flashcards`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flashcards []Flashcard

	for rows.Next() {
		var flashcard Flashcard
		err := rows.Scan(&flashcard.ID, &flashcard.Front, &flashcard.Back, &flashcard.CreatedAt)
		if err != nil {
			return nil, errors.New("error while scaning parameters")
		}
		flashcards = append(flashcards, flashcard)
	}

	return flashcards, nil
}