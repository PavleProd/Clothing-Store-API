package db

import (
	"online_store_api/src/util"
	"reflect"
	"testing"
)

var testCasesBuildSelectQuery = []struct {
	name      string
	data      util.DataRecord
	tableName string
	expected  PreparedQuery
}{
	{
		"NoParams_Valid",
		util.DataRecord{},
		"products",
		*NewPreparedQuery(
			"SELECT * FROM products",
			[]any{},
		),
	},
	{
		"OneParam_Valid",
		util.DataRecord{"name": "T Shirt"},
		"products",
		*NewPreparedQuery(
			"SELECT * FROM products WHERE name = $1",
			[]any{"T Shirt"},
		),
	},
}

func TestBuildSelectQuery(t *testing.T) {
	for _, testcase := range testCasesBuildSelectQuery {
		t.Run(testcase.name, func(t *testing.T) {
			expected := testcase.expected

			got := BuildSelectQuery(testcase.data, testcase.tableName)

			if !reflect.DeepEqual(got, expected) {
				t.Errorf("got %v expected %v", got, expected)
			}
		})
	}
}
