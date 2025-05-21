package flashcards

import (
	"encoding/json"
	"errors"
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