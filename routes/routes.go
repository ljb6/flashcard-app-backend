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

	flashcardGroup := router.Group("/flashcards")
	flashcardGroup.POST("/create-flashcard", flashcardsHandler.CreateFlashcardHandler)
	flashcardGroup.POST("/delete-flashcard", flashcardsHandler.DeleteFlashcardByIDHandler)
	flashcardGroup.PATCH("/update-flashcards", flashcardsHandler.EditFlashcardByIDHandler)
	flashcardGroup.GET("/get-flashcards", flashcardsHandler.GetFlashcardsHandler)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
