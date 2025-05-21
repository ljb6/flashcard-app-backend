package flashcards

import "database/sql"

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