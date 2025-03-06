package db

import (
	"fmt"
	"online_store_api/src/util"
	"strings"
)

// SELECT * FROM TABLE WHERE param1 = value1 AND param2 = value2
func BuildSelectQuery(data util.DataRecord, tableName string) PreparedQuery {
	var queryBuilder strings.Builder
	var values = []any{}

	queryBuilder.WriteString("SELECT * FROM ")
	queryBuilder.WriteString(tableName)

	var i int = 0
	for key, value := range data {
		if i == 0 {
			queryBuilder.WriteString(" WHERE ")
		} else {
			queryBuilder.WriteString(" AND ")
		}

		// PostgreSQL uses placeholders $1, $2, $3, ...
		var placeholder string = fmt.Sprintf("$%v", i+1)
		queryBuilder.WriteString(key)
		queryBuilder.WriteString(" = ")
		queryBuilder.WriteString(placeholder)

		values = append(values, value)
		i++
	}

	return *NewPreparedQuery(queryBuilder.String(), values)
}

// INSERT INTO TABLE (field1, field2)
// VALUES (value1, value2)
func BuildInsertQuery(data util.DataRecord, tableName string) PreparedQuery {
	var queryBuilder strings.Builder

	queryBuilder.WriteString("INSERT INTO ")
	queryBuilder.WriteString(tableName)

	var keys = make([]string, 0, len(data))
	var values = make([]any, 0, len(data))
	for key, value := range data {
		keys = append(keys, key)
		values = append(values, value)
	}

	queryBuilder.WriteString("(" + strings.Join(keys, ",") + ")")

	// PostgreSQL uses placeholders $1, $2, $3, ...
	queryBuilder.WriteString(" VALUES(")
	for i := range len(data) {
		if i != 0 {
			queryBuilder.WriteString(", ")
		}
		var placeholder string = fmt.Sprintf("$%v", i+1)
		queryBuilder.WriteString(placeholder)
	}
	queryBuilder.WriteString(")")

	return *NewPreparedQuery(queryBuilder.String(), values)
}
