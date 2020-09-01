package utils

import "testing"

func AssertEquals(t *testing.T, expected interface{}, actual interface{}, message string) {
	if expected != actual {
		t.Error(message, "-> expecting:", expected, "got:", actual)
	}
}

func AssertTrue(t *testing.T, condition bool, message string) {
	if !condition {
		t.Error(message)
	}
}
