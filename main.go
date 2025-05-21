package main

import "github.com/ljb6/flashcard-app-backend/database"

func main() {
	db := database.ConnectDB()
	defer db.Close()
}