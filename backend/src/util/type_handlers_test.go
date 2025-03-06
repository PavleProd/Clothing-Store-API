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

var testCasesIsDefaultOrZeroValueExcludingBool = []struct {
	name     string
	value    any
	expected bool
}{
	{"int default", 0, true},
	{"int non-default", 1, false},
	{"float32 default", 0.0, true},
	{"float32 non-default", 1.0, false},
	{"string default", "", true},
	{"string non-default", "1", false},
	{"bool default", false, false},
	{"bool non-default", true, false},
}

func TestIsDefaultOrZeroValueExcludingBool(t *testing.T) {
	for _, testcase := range testCasesIsDefaultOrZeroValueExcludingBool {
		t.Run(testcase.name, func(t *testing.T) {
			var expected bool = testcase.expected

			var got bool = IsDefaultOrZeroValueExcludingBool(testcase.value)

			if got != expected {
				t.Errorf("excpected %v, got %v", expected, got)
			}
		})
	}
}
