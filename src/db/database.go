package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectToDatabase(databaseName string) *sql.DB {
	var connectionString = fmt.Sprintf("user=%v dbname=%v password=%v sslmode=disable", Admin.Username, databaseName, Admin.Password)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic(err)
	}

	return db
}
