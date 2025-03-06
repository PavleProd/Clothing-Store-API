package db

import (
	"online_store_api/src/util"
	"strings"
)

// SELECT * FROM TABLE WHERE param1 = value1 AND param2 = value2
func BuildSelectQuery(data util.DataRecord, tableName string) string {
	var queryBuilder strings.Builder

	queryBuilder.WriteString("SELECT * FROM ")
	queryBuilder.WriteString(tableName)

	var firstParameter = true
	for key, value := range data {
		if firstParameter {
			firstParameter = false
			queryBuilder.WriteString(" WHERE ")
		} else {
			queryBuilder.WriteString(" AND ")
		}

		queryBuilder.WriteString(key)
		queryBuilder.WriteString(" = ")
		queryBuilder.WriteString("'" + value + "'")
	}

	return queryBuilder.String()
}

// INSERT INTO TABLE (field1, field2)
// VALUES (value1, value2)
func BuildInsertQuery(data util.DataRecord, tableName string) string {
	var queryBuilder strings.Builder

	queryBuilder.WriteString("INSERT INTO ")
	queryBuilder.WriteString(tableName)

	var keys = make([]string, 0, len(data))
	var values = make([]string, 0, len(data))
	for key, value := range data {
		keys = append(keys, key)
		values = append(values, "'"+value+"'")
	}

	queryBuilder.WriteString("(" + strings.Join(keys, ",") + ")")
	queryBuilder.WriteString(" VALUES(" + strings.Join(values, ",") + ")")

	return queryBuilder.String()
}
