package db

import (
	"errors"
	"fmt"
	"online_store_api/src/util"
	"strings"
)

// SELECT * FROM TABLE WHERE param1 = value1 AND param2 = value2
func BuildReadQuery[T any](model T, tableName string) string {
	var queryBuilder strings.Builder

	queryBuilder.WriteString("SELECT * FROM ")
	queryBuilder.WriteString(tableName)

	var slicedFields = util.GetModelSlicedFields(model)
	var firstParameter bool = true
	for i := range slicedFields {
		if util.IsDefaultOrZeroValueExcludingBool(slicedFields[i].Value) {
			continue
		}

		if firstParameter {
			firstParameter = false
			queryBuilder.WriteString(" WHERE ")
		} else {
			queryBuilder.WriteString(" AND ")
		}

		var valueAsString string = fmt.Sprint("'", slicedFields[i].Value, "'")

		queryBuilder.WriteString(slicedFields[i].Tag)
		queryBuilder.WriteString("=")
		queryBuilder.WriteString(valueAsString)
	}

	return queryBuilder.String()
}

// INSERT INTO TABLE (field1, field2)
// VALUES (value1, value2)
func BuildInsertQuery[T any](model T, tableName string) (string, error) {
	var queryBuilder strings.Builder

	var slicedFields = util.GetModelSlicedFields(model)
	for i := range slicedFields {
		if util.IsDefaultOrZeroValueExcludingBool(slicedFields[i].Value) {
			var errorMessage = fmt.Sprintf("field %v not defined", slicedFields[i].Name)
			return "", errors.New(errorMessage)
		}
	}

	queryBuilder.WriteString("INSERT INTO ")
	queryBuilder.WriteString(tableName)

	queryBuilder.WriteString(" (")
	for i := range slicedFields {
		if i != 0 {
			queryBuilder.WriteString(", ")
		}
		queryBuilder.WriteString(slicedFields[i].Tag)
	}
	queryBuilder.WriteString(")")

	queryBuilder.WriteString(" VALUES (")
	for i := range slicedFields {
		if i != 0 {
			queryBuilder.WriteString(", ")
		}

		var strValue = fmt.Sprint("'", slicedFields[i].Value, "'")
		queryBuilder.WriteString(strValue)
	}
	queryBuilder.WriteString(")")

	return queryBuilder.String(), nil
}
