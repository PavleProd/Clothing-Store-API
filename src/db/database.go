package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type ConnectionManager struct {
	DatabaseName string
	instance     *sql.DB
}

func (manager *ConnectionManager) Connect() {
	var connectionString = fmt.Sprintf("user=%v dbname=%v password=%v sslmode=disable", Admin.Username, manager.DatabaseName, Admin.Password)

	instance, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic(err)
	}

	manager.instance = instance
}

func (manager *ConnectionManager) Close() {
	manager.instance.Close()
}
