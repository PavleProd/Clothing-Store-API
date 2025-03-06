package db

import (
	"online_store_api/src/model"
)

var testCasesBuildReadQuery = []struct {
	name       string
	model      any
	tableName  string
	expected   string
	shouldFail bool
}{
	{"all default", model.Product{}, "products", "SELECT * FROM products", false},
}

// TODO: fix tests when code is refactored
// func TestBuildReadQuery(t *testing.T) {
// 	for _, testcase := range testCasesBuildReadQuery {
// 		t.Run(testcase.name, func(t *testing.T) {
// 			expected := testcase.expected

// 			got := BuildSelectQueryFromModel(testcase.model, testcase.tableName)

// 			// when testcase should pass then it's wrong that got is expected. Same logic when testcase should pass
// 			if testcase.shouldFail != (got == expected) {
// 				t.Errorf("got: %v, expected: %v", got, expected)
// 			}
// 		})
// 	}
// }
