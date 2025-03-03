package db

import (
	"database/sql"
	"online_store_api/src/util"
)

func ConvertToDataSet(resultSet *sql.Rows) (util.DataSet, error) {
	var result util.DataSet

	columnNames, err := resultSet.Columns()
	if err != nil {
		return result, err
	}

	var numColumns = len(columnNames)

	// Scan method needs "numColumns" number of pointers (rowPointers)
	// so it can write result into referenced variables (rowValues)
	var rowValues = make([]string, numColumns)
	var rowPointers = make([]any, numColumns)
	for i := range numColumns {
		rowPointers[i] = &rowValues[i]
	}

	for resultSet.Next() {
		resultSet.Scan(rowPointers...)

		var dataRecord = make(util.DataRecord)
		for i := range numColumns {
			dataRecord[columnNames[i]] = rowValues[i]
		}
		result = append(result, dataRecord)
	}

	return result, nil
}
