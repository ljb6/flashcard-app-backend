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
	flashcardGroup.POST("/create", flashcardsHandler.CreateFlashcardHandler)
	flashcardGroup.POST("/delete", flashcardsHandler.DeleteFlashcardByIDHandler)
	flashcardGroup.POST("/deleteall", flashcardsHandler.DeleteAllFlashcardsHandler)
	flashcardGroup.PATCH("/update", flashcardsHandler.EditFlashcardByIDHandler)
	flashcardGroup.PATCH("/update-stats", flashcardsHandler.UpdateFlashcardFieldsByIDHandler)
	// usando POST por conta do body
	flashcardGroup.POST("/get", flashcardsHandler.GetFlashcardsHandler)
	flashcardGroup.GET("/get-due", flashcardsHandler.GetDueFlashcardsHandler)


	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}
