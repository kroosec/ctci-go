package main

import "testing"

func TestIsOnlyOneEdit(t *testing.T) {
	testInputs := []struct {
		original string
		edit     string
		want     bool
	}{
		// Replace cases.
		{"a", "a", false},
		{"a", "b", true},
		{"abcdef", "abcref", true},
		{"abcdef", "abcrer", false},
		// Removal cases.
		{"a", "", true},
		{"abc", "ab", true},
		{"abc", "ac", true},
		{"abcc", "abc", true},
		{"abcc", "bcc", true},
		{"abcc", "bdc", false},
		{"abcc", "dcc", false},
		{"abcc", "bcd", false},
		{"abcc", "abd", false},
		// Insert cases
		{"", "a", true},
		{"ab", "abc", true},
		{"ac", "abc", true},
		{"abc", "abcc", true},
		{"bcc", "abcc", true},
		{"bdc", "abcc", false},
		{"dcc", "abcc", false},
		{"bcd", "abcc", false},
		{"abd", "abcc", false},
		// Other cases
		{"aa", "", false},
		{"a", "aba", false},
	}
	for _, input := range testInputs {
		got := IsOnlyOneEdit(input.original, input.edit)
		AssertInputs(t, input, got, input.want)
	}
}
