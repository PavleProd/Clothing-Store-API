package db

import (
	"database/sql"
	"online_store_api/src/util"

	_ "github.com/lib/pq"
)

type ConnectionManager struct {
	instance *sql.DB
}

func (manager *ConnectionManager) Connect(connectionURL string) error {
	instance, err := sql.Open("postgres", connectionURL)
	if err != nil {
		return err
	}

	manager.instance = instance
	return nil
}

func (manager *ConnectionManager) Close() {
	manager.instance.Close()
}

func (manager *ConnectionManager) Read(query string) (util.DataSet, error) {
	var result util.DataSet

	resultSet, err := manager.instance.Query(query)
	if err != nil {
		return result, err
	}

	return ConvertToDataSet(resultSet)
}
