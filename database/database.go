package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB_PASSWORD string

func GetDBPassword() {
	DB_PASSWORD = os.Getenv("PASSWORD")
}

const (
	host     = "dpg-d0nq3ps9c44c73evbk0g-a.oregon-postgres.render.com"
	port     = 5432
	user     = "postgres_flashapp_user"
	dbname   = "postgres_flashapp"
)

func ConnectDB() (*sql.DB) {
	postgreSqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, DB_PASSWORD, dbname)
	db, err := sql.Open("postgres", postgreSqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB: " + dbname)
	
	return db
}