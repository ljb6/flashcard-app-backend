package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	//"github.com/ljb6/flashcard-app-backend/internal/flashcards"
)

func InitializeRoutes(router *gin.Engine, db *sql.DB) {

	// flashcards
	// flashcardsRepo := flashcards.NewFlashcardRepository(db)
	// flashcardsService := flashcards.NewFlashcardService(flashcardsRepo)
	// flashcardsHandler := flashcards.NewFlashcardHandlers(flashcardsService)

	//flashcardGroup := router.Group("/flashcards")

	router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

}