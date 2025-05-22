package flashcards

import "time"

type Flashcard struct {
	ID               int        `json:"id"`
	Front            string     `json:"front"`
	Back             string     `json:"back"`
	CreatedAt        time.Time  `json:"created_at"`
	LastReview       *time.Time `json:"last_review"`
	ReviewStage      int        `json:"review_stage"`
	CorrectAnswers   int        `json:"correct_answers"`
	IncorrectAnswers int        `json:"incorrect_answers"`
}

var ReviewIntervals = []int{0, 1, 3, 7, 15, 30}

type GetFlashcardsReq struct {
	Quantity int    `json:"quantity"`
}

type UpdateFlashcardFieldsReq struct {
	ID      int  `json:"id"`
	Correct bool `json:"correct"`
}
