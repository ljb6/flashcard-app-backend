package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "lucca"
	password = "1234"
	dbname   = "postgres"
)

func ConnectDB() (*sql.DB) {
	postgreSqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgreSqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB: " + dbname)
	
	return db
}