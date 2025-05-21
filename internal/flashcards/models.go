package flashcards

import "time"

type Flashcard struct {
	ID        int       `jsond:"id"`
	Front     string    `jsond:"front"`
	Back      string    `jsond:"back"`
	CreatedAt time.Time `json:"created_at"`
}

type GetFlashcardsReq struct {
	ReqType  string `jsond:"req_type"`
	Quantity int    `jsond:"quantity"`
}
