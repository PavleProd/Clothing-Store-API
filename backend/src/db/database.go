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

func (manager *DatabaseManager) Read(preparedQuery PreparedQuery) (util.DataSet, error) {
	statement, err := manager.instance.Prepare(preparedQuery.query)
	if err != nil {
		return util.DataSet{}, err
	}
	defer statement.Close()

	resultSet, err := statement.Query(preparedQuery.values...)
	if err != nil {
		return util.DataSet{}, err
	}

	return ConvertToDataSet(resultSet)
}

func (manager *DatabaseManager) Write(preparedQuery PreparedQuery) error {
	statement, err := manager.instance.Prepare(preparedQuery.query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(preparedQuery.values...)
	return err
}
