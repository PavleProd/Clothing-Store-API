package db

import (
	"database/sql"
	"online_store_api/src/util"

	_ "github.com/lib/pq"
)

type DatabaseManager struct {
	instance *sql.DB
}

func (manager *DatabaseManager) Connect(connectionURL string) error {
	instance, err := sql.Open("postgres", connectionURL)
	if err != nil {
		return err
	}

	manager.instance = instance
	return nil
}

func (manager *DatabaseManager) Close() {
	manager.instance.Close()
}

func (manager *DatabaseManager) Read(query string) (util.DataSet, error) {
	var result util.DataSet

	resultSet, err := manager.instance.Query(query)
	if err != nil {
		return result, err
	}

	return ConvertToDataSet(resultSet)
}

func (manager *DatabaseManager) Write(query string) error {
	_, err := manager.instance.Exec(query)
	return err
}
