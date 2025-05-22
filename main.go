package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/ljb6/flashcard-app-backend/database"
	"github.com/ljb6/flashcard-app-backend/internal/flashcards"
	router "github.com/ljb6/flashcard-app-backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("erro ao carregar keys")
	}

	flashcards.GetKeys()

	db := database.ConnectDB()
	defer db.Close()

	database.CreateTables(db)

	router.InitializeServer(db)
}