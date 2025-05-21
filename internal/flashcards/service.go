package flashcards

type FlashcardService struct {
	repository *FlashcardRepository
}

func NewFlashcardService(repository *FlashcardRepository) *FlashcardService {
	return &FlashcardService{repository: repository}
}