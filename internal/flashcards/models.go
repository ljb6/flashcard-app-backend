package flashcards

import "time"

type Flashcard struct {
	ID        int       `json:"id"`
	Front     string    `json:"front"`
	Back      string    `json:"back"`
	CreatedAt time.Time `json:"created_at"`
}

type GetFlashcardsReq struct {
	ReqType  string `json:"req_type"`
	Quantity int    `json:"quantity"`
}
