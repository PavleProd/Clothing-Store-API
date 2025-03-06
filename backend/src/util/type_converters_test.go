package util

import (
	"reflect"
	"testing"
)

var testCasesConvertFromString = []struct {
	name          string
	value         string
	expected      any
	errorExpected bool
}{
	{"ConvertToBool_Valid", "false", false, false},
	{"ConvertToNegativeInt_Valid", "-10", -10, false},
	{"ConverToPositiveInt_Valid", "10", 10, false},
	{"ConvertToString_Valid", "test", "test", false},
	{"ConvertToFloat_Valid", "3.14", 3.14, false},
	{"ConvertToInt_InvalidFloatValue", "3.14", 3, true},
}

func TestConvertFromString(t *testing.T) {
	for _, testcase := range testCasesConvertFromString {
		t.Run(testcase.name, func(t *testing.T) {
			var expected any = testcase.expected

			got, err := ConvertFromString(testcase.value, reflect.TypeOf(expected))

			if testcase.errorExpected {
				if err == nil {
					t.Errorf("error expected, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("error message not expected: %v", err.Error())
				} else if got.Interface() != expected {
					t.Errorf("expected %v, got %v", expected, got)
				}
			}
		})
	}
}
