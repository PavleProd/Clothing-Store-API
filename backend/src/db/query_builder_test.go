package db

import (
	"online_store_api/src/util"
	"reflect"
	"testing"
)

var testCasesBuildSelectQuery = []struct {
	name       string
	data       util.DataRecord
	tableName  string
	expected   PreparedQuery
	shouldPass bool
}{
	{
		"NoParams_Valid",
		util.DataRecord{},
		"products",
		*NewPreparedQuery(
			"SELECT * FROM products",
			[]any{},
		),
		true,
	},
	{
		"OneParam_Valid",
		util.DataRecord{"name": "T Shirt"},
		"products",
		*NewPreparedQuery(
			"SELECT * FROM products WHERE name = $1",
			[]any{"T Shirt"},
		),
		true,
	},
	{
		"OneParam_InvalidValuesSlice",
		util.DataRecord{"name": "T Shirt"},
		"products",
		*NewPreparedQuery(
			"SELECT * FROM products WHERE name = $1",
			[]any{},
		),
		false,
	},
	{
		"OneParam_InvalidQuery",
		util.DataRecord{"name": "T Shirt"},
		"products",
		*NewPreparedQuery(
			"SELECT * FROM products WHERE name = $2",
			[]any{"T Shirt"},
		),
		false,
	},
}

func TestBuildSelectQuery(t *testing.T) {
	for _, testcase := range testCasesBuildSelectQuery {
		t.Run(testcase.name, func(t *testing.T) {
			expected := testcase.expected

			got := BuildSelectQuery(testcase.data, testcase.tableName)

			var isEqual bool = reflect.DeepEqual(got, expected)
			if isEqual != testcase.shouldPass {
				t.Errorf("got %v expected %v", got, expected)
			}
		})
	}
}
