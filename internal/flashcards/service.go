package flashcards

import "errors"

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