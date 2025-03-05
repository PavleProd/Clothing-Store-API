package db

import (
	"errors"
	"fmt"
	"online_store_api/src/util"
	"reflect"
	"strings"
)

func BuildReadQuery[T any](model T, tableName string) string {
	var queryBuilder strings.Builder

	queryBuilder.WriteString("SELECT * FROM ")
	queryBuilder.WriteString(tableName)

	var firstParameter bool = true
	var reflectedModelValue = reflect.ValueOf(&model).Elem()
	var reflectedModelType = reflectedModelValue.Type()
	for i := range reflectedModelValue.NumField() {
		var reflectedFieldValue = reflectedModelValue.Field(i).Interface()
		if util.IsDefaultOrZeroValueExcludingBool(reflectedFieldValue) {
			continue
		}

		if firstParameter {
			firstParameter = false
			queryBuilder.WriteString(" WHERE ")
		} else {
			queryBuilder.WriteString(" AND ")
		}

		var reflectedFieldType = reflectedModelType.Field(i)
		var reflectedFieldName = reflectedFieldType.Tag.Get(util.JSON_TAG)
		var valueAsString string = fmt.Sprint(reflectedFieldValue)

		queryBuilder.WriteString(reflectedFieldName)
		queryBuilder.WriteString("=")
		queryBuilder.WriteString("'")
		queryBuilder.WriteString(valueAsString)
		queryBuilder.WriteString("'")
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
