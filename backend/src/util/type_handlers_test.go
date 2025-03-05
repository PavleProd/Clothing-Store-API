package util

import "testing"

var testcases = []struct {
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
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			var expected bool = testcase.expected

			var got bool = IsDefaultOrZeroValueExcludingBool(testcase.value)

			if got != expected {
				t.Errorf("excpected %v, got %v", expected, got)
			}
		})
	}
}
