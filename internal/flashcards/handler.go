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

	var req Flashcard

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

func (h *FlashcardHandler) GetAllFlashcardsHandler(c *gin.Context) {

	jsonFlashcards, err := h.service.GetAllFlashcards()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "impossible to get all flashcards"})
		return
	}

	c.Data(http.StatusOK, "application/json", jsonFlashcards)
}

func (h *FlashcardHandler) GetXRandomFlashcardsHandler(c *gin.Context) {

	var req GetFlashcardsReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request format error"})
		return
	}

	jsonFlashcards, err := h.service.GetXFlashcards(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "impossible to get flashcards"})
		return
	}

	c.Data(http.StatusOK, "application/json", jsonFlashcards)
}

func (h *FlashcardHandler) GetXRandomFlashcardsByErrorHandler(c *gin.Context) {

	var req GetFlashcardsReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request format error"})
		return
	}

	jsonFlashcards, err := h.service.GetXFlashcardsByError(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "impossible to get flashcards"})
		return
	}

	c.Data(http.StatusOK, "application/json", jsonFlashcards)
}

func (h *FlashcardHandler) DeleteFlashcardByIDHandler(c *gin.Context) {

	var flashcard Flashcard
	err := c.ShouldBindJSON(&flashcard)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid flashcard"})
		return
	}

	err = h.service.DeleteFlashcardByID(flashcard.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "impossible to delete flashcard"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flashcard deleted with success"})
}

func (h *FlashcardHandler) EditFlashcardByIDHandler(c *gin.Context) {

	var req Flashcard

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameters"})
		return
	}

	err = h.service.EditFlashcardByID(req.ID, req.Front, req.Back)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flashcard updated with success"})
}

func (h *FlashcardHandler) DeleteAllFlashcardsHandler(c *gin.Context) {

	err := h.service.DeleteAllFlashcards()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flashcard deleted with success"})
}

func (h *FlashcardHandler) UpdateFlashcardFieldsByIDHandler(c *gin.Context) {

	var req UpdateFlashcardFieldsReq

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid parameters"})
		return
	}

	err = h.service.UpdateFlashcardFields(req.ID, req.Correct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "flashcard updated with success"})
}

func (h *FlashcardHandler) GetDueFlashcardsHandler(c *gin.Context) {

	jsonFlashcards, err := h.service.GetDueFlashcards()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "impossible to get due flashcards"})
		return
	}

	c.Data(http.StatusOK, "application/json", jsonFlashcards)
}