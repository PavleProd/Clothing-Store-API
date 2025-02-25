package db

import (
	"database/sql"
	"fmt"
	"online_store_api/src/util"

	_ "github.com/lib/pq"
)

type ConnectionManager struct {
	DatabaseName string
	instance     *sql.DB
}

func (manager *ConnectionManager) Connect() error {
	var connectionString = fmt.Sprintf("user=%v dbname=%v password=%v sslmode=disable", Admin.Username, manager.DatabaseName, Admin.Password)

	instance, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}

	manager.instance = instance
	return nil
}

func (manager *ConnectionManager) Close() {
	manager.instance.Close()
}

func (manager *ConnectionManager) Query(query string) (util.DataSet, error) {
	var result util.DataSet

	resultSet, err := manager.instance.Query(query)
	if err != nil {
		return result, err
	}

	return ConvertToDataSet(resultSet)
}
