package util

import (
	"reflect"
	"testing"
)

var testCasesConvertFromString = []struct {
	name       string
	value      string
	expected   any
	shouldPass bool
}{
	{"ConvertToBool_Valid", "false", false, true},
	{"ConvertToNegativeInt_Valid", "-10", -10, true},
	{"ConverToPositiveInt_Valid", "10", 10, true},
	{"ConvertToString_Valid", "test", "test", true},
	{"ConvertToFloat_Valid", "3.14", 3.14, true},
	{"ConvertToInt_InvalidFloatValue", "3.14", 3, false},
}

func TestConvertFromString(t *testing.T) {
	for _, testcase := range testCasesConvertFromString {
		t.Run(testcase.name, func(t *testing.T) {
			var expected any = testcase.expected

			got, err := ConvertFromString(testcase.value, reflect.TypeOf(expected))

			if testcase.shouldPass {
				if err != nil {
					t.Errorf("error message not expected: %v", err.Error())
				} else if got.Interface() != expected {
					t.Errorf("expected %v, got %v", expected, got)
				}
			} else {
				if err == nil {
					t.Errorf("error expected, but got nil")
				}
			}
		})
	}
}
