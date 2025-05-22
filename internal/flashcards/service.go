package flashcards

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type FlashcardService struct {
	repository *FlashcardRepository
}

func NewFlashcardService(repository *FlashcardRepository) *FlashcardService {
	return &FlashcardService{repository: repository}
}

func (s *FlashcardService) CreateFlashcard(front, back string) error {
	if len(front) > 250 || len(back) > 250 {
		return errors.New("flashcard content exceeds maximum length of 250 characters")
	}

	err := s.repository.CreateFlashcard(front, back)
	if err != nil {
		return err
	}

	return nil
}

func (s *FlashcardService) GetFlashcards(req GetFlashcardsReq) ([]byte, error) {
	flashcards, err := s.repository.GetFlashcards(req)
	if err != nil {
		return nil, err
	}

	jsonFlashcards, err := json.Marshal(flashcards)
	if err != nil {
		return nil, errors.New("error in marshal")
	}

	return jsonFlashcards, nil
}

func (s *FlashcardService) DeleteFlashcardByID(id int) error {
	err := s.repository.DeleteFlashcardByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *FlashcardService) EditFlashcardByID(id int, front, back string) error {
	if len(front) > 250 || len(back) > 250 {
		return errors.New("flashcard content exceeds maximum length of 250 characters")
	}

	err := s.repository.EditFlashcardByID(id, front, back)
	if err != nil {
		return err
	}

	return nil
}

func (s *FlashcardService) DeleteAllFlashcards() error {
	err := s.repository.DeleteAllFlashcards()
	if err != nil {
		return err
	}
	return nil
}

func (s *FlashcardService) UpdateFlashcardFields(id int, correct bool) error {
	card, err := s.repository.GetFlashcardByID(id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	now := time.Now()
	card.LastReview = &now

	if correct {
		if card.ReviewStage < len(ReviewIntervals)-1 {
			card.ReviewStage++
		}
		card.CorrectAnswers++
	} else {
		card.IncorrectAnswers++
		card.ReviewStage = 0
	}

	err = s.repository.UpdateFlashcardFields(card)
	if err != nil {
		return err
	}
	return nil
}
