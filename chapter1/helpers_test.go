package main

import "testing"

func AssertBool(t *testing.T, input interface{}, got, want bool) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v want %v, given '%+v'", got, want, input)
	}
}

