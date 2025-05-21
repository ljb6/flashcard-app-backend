package flashcards

type FlashcardHandler struct {
	service *FlashcardService
}

func NewFlashcardHandlers(service *FlashcardService) *FlashcardHandler {
	return &FlashcardHandler{service: service}
}