package db

import (
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
		if util.IsDefaultValue(reflectedFieldValue) {
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
