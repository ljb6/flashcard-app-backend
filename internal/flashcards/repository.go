package flashcards

import (
	"database/sql"
	"errors"
	"fmt"
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

func (r *FlashcardRepository) GetAllFlashcards() ([]Flashcard, error) {
	query := "SELECT id, front, back, created_at, last_review, review_stage, correct_answers, incorrect_answers FROM flashcards"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flashcards []Flashcard

	for rows.Next() {
		var card Flashcard
		err := rows.Scan(&card.ID, &card.Front, &card.Back, &card.CreatedAt, &card.LastReview,
			&card.ReviewStage, &card.CorrectAnswers, &card.IncorrectAnswers)
		if err != nil {
			return nil, errors.New("error while scanning parameters")
		}
		flashcards = append(flashcards, card)
	}

	return flashcards, nil
}

func (r *FlashcardRepository) GetXFlashcards(req GetFlashcardsReq) ([]Flashcard, error) {
	query := fmt.Sprintf("SELECT id, front, back, created_at, last_review, review_stage, correct_answers, incorrect_answers FROM flashcards ORDER BY RANDOM() LIMIT %d", req.Quantity)

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flashcards []Flashcard

	for rows.Next() {
		var card Flashcard
		err := rows.Scan(&card.ID, &card.Front, &card.Back, &card.CreatedAt, &card.LastReview,
			&card.ReviewStage, &card.CorrectAnswers, &card.IncorrectAnswers)
		if err != nil {
			return nil, errors.New("error while scaning parameters")
		}
		flashcards = append(flashcards, card)
	}

	return flashcards, nil
}

func (r *FlashcardRepository) GetXFlashcardsByError(req GetFlashcardsReq) ([]Flashcard, error) {
	query := fmt.Sprintf(`SELECT id, front, back, created_at, last_review, review_stage, correct_answers,
							incorrect_answers FROM flashcards WHERE (incorrect_answers + correct_answers) > 0
							ORDER BY CAST(incorrect_answers AS FLOAT) / (incorrect_answers + correct_answers)
							DESC LIMIT %d`, req.Quantity)
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var flashcards []Flashcard

	for rows.Next() {
		var card Flashcard
		err := rows.Scan(&card.ID, &card.Front, &card.Back, &card.CreatedAt, &card.LastReview,
			&card.ReviewStage, &card.CorrectAnswers, &card.IncorrectAnswers)
		if err != nil {
			return nil, errors.New("error while scaning parameters")
		}
		flashcards = append(flashcards, card)
	}

	return flashcards, nil
}

func (r *FlashcardRepository) GetFlashcardByID(id int) (Flashcard, error) {
	query := "SELECT id, front, back, created_at, last_review, review_stage, correct_answers, incorrect_answers FROM flashcards WHERE id = $1"

	var flashcard Flashcard
	err := r.DB.QueryRow(query, id).Scan(&flashcard.ID, &flashcard.Front, &flashcard.Back, &flashcard.CreatedAt, &flashcard.LastReview, &flashcard.ReviewStage, &flashcard.CorrectAnswers, &flashcard.IncorrectAnswers)
	if err != nil {
		return Flashcard{}, err
	}

	return flashcard, nil
}

func (r *FlashcardRepository) DeleteFlashcardByID(id int) error {
	query := `DELETE FROM flashcards WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *FlashcardRepository) EditFlashcardByID(id int, front, back string) error {
	query := `UPDATE flashcards SET front = $1, back = $2 WHERE id = $3`
	_, err := r.DB.Exec(query, front, back, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *FlashcardRepository) DeleteAllFlashcards() error {
	query := `DELETE FROM flashcards`
	_, err := r.DB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (r *FlashcardRepository) UpdateFlashcardFields(card Flashcard) error {
	query := `UPDATE flashcards SET last_review = $1, review_stage = $2, correct_answers = $3, incorrect_answers = $4 WHERE id = $5`
	_, err := r.DB.Exec(query, card.LastReview, card.ReviewStage, card.CorrectAnswers, card.IncorrectAnswers, card.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *FlashcardRepository) GetDueFlashcards() ([]Flashcard, error) {
	var dueCards []Flashcard

	query1 := `SELECT id, front, back, created_at, last_review, review_stage, correct_answers, incorrect_answers 
               FROM flashcards 
               WHERE last_review IS NULL`

	rows1, err := r.DB.Query(query1)
	if err != nil {
		return nil, err
	}
	defer rows1.Close()

	for rows1.Next() {
		var card Flashcard
		err := rows1.Scan(&card.ID, &card.Front, &card.Back, &card.CreatedAt, &card.LastReview,
			&card.ReviewStage, &card.CorrectAnswers, &card.IncorrectAnswers)
		if err != nil {
			return nil, err
		}
		dueCards = append(dueCards, card)
	}

	for stage, interval := range ReviewIntervals {
		query := `SELECT id, front, back, created_at, last_review, review_stage, correct_answers, incorrect_answers 
                 FROM flashcards 
                 WHERE last_review IS NOT NULL 
                   AND review_stage = $1
                   AND NOW() >= last_review + (INTERVAL '1 day' * $2)`

		rows, err := r.DB.Query(query, stage, interval)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var card Flashcard
			err := rows.Scan(&card.ID, &card.Front, &card.Back, &card.CreatedAt, &card.LastReview,
				&card.ReviewStage, &card.CorrectAnswers, &card.IncorrectAnswers)
			if err != nil {
				return nil, err
			}
			dueCards = append(dueCards, card)
		}
	}

	return dueCards, nil
}
