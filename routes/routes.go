package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ljb6/flashcard-app-backend/internal/flashcards"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {

	// flashcards
	flashcardsRepo := flashcards.NewFlashcardRepository(db)
	flashcardsService := flashcards.NewFlashcardService(flashcardsRepo)
	flashcardsHandler := flashcards.NewFlashcardHandlers(flashcardsService)

	// localhost:8080/flashcards/...
	flashcardGroup := router.Group("/flashcards")

	// POST
	flashcardGroup.POST("/create", flashcardsHandler.CreateFlashcardHandler)
	flashcardGroup.POST("/delete", flashcardsHandler.DeleteFlashcardByIDHandler)
	flashcardGroup.POST("/delete-all", flashcardsHandler.DeleteAllFlashcardsHandler)
	flashcardGroup.POST("/get-random", flashcardsHandler.GetXRandomFlashcards) // post por conta do body

	// PATCH 
	flashcardGroup.PATCH("/update", flashcardsHandler.EditFlashcardByIDHandler)
	flashcardGroup.PATCH("/update-stats", flashcardsHandler.UpdateFlashcardFieldsByIDHandler)
	
	// GET
	flashcardGroup.GET("/get-all", flashcardsHandler.GetAllFlashcardsHandler)
	flashcardGroup.GET("/get-due", flashcardsHandler.GetDueFlashcardsHandler)
}
