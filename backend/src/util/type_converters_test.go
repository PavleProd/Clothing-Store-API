package util

import (
	"reflect"
	"testing"
)

var testCasesConvertFromString = []struct {
	name       string
	value      string
	expected   any
	shouldFail bool
}{
	{"convert to bool", "false", false, false},
	{"convert to negative int", "-10", -10, false},
	{"convert to positive int", "10", 10, false},
	{"convert to string", "test", "test", false},
	{"convert to float", "3.14", 3.14, false},
	{"invalid target int, value float", "3.14", 3, true},
}

func TestConvertFromString(t *testing.T) {
	for _, testcase := range testCasesConvertFromString {
		t.Run(testcase.name, func(t *testing.T) {
			var expected any = testcase.expected

			got, err := ConvertFromString(testcase.value, reflect.TypeOf(expected))

			if testcase.shouldFail {
				if err == nil {
					t.Errorf("error expected, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("error message not expected: %v", err.Error())
				} else if got.Interface() != expected {
					t.Errorf("excpected %v, got %v", expected, got)
				}
			}
		})
	}
}
