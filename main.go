package main

import (
	"github.com/ljb6/flashcard-app-backend/database"
	router "github.com/ljb6/flashcard-app-backend/routes"
)

func main() {
	db := database.ConnectDB()
	defer db.Close()

	router.InitializeServer(db)
}