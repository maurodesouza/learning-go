package utils

import "testing"

func AssertEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func Assert(t *testing.T, cond bool, msg string) {
	t.Helper()

	if !cond {
		t.Fatal(msg)
	}
}
