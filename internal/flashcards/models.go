package flashcards

import "time"

type Flashcard struct {
	ID             int        `json:"id"`
	Front          string     `json:"front"`
	Back           string     `json:"back"`
	CreatedAt      time.Time  `json:"created_at"`
	LastReview     *time.Time `json:"last_review"`
	ReviewStage    int        `json:"review_stage"`
	CorrectAnswers int        `json:"correct_answers"`
	Tries          int        `json:"tries"`
}

type GetFlashcardsReq struct {
	ReqType  string `json:"req_type"`
	Quantity int    `json:"quantity"`
}
