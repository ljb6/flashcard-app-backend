package flashcards

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlashcardHandler struct {
	service *FlashcardService
}

func NewFlashcardHandlers(service *FlashcardService) *FlashcardHandler {
	return &FlashcardHandler{service: service}
}

func (h *FlashcardHandler) CreateFlashcardHandler(c *gin.Context) {

	var req Flashcard;

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameters"})
		return
	}

	err = h.service.CreateFlashcard(req.Front, req.Back)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flashcard added with success"})
}

func (h *FlashcardHandler) GetFlashcardsHandler(c *gin.Context) {
	jsonFlashcards, err := h.service.GetFlashcards()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "impossible to get flashcards"})
		return
	}

	c.Data(http.StatusOK, "application/json", jsonFlashcards)
}