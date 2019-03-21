package main

import "testing"

func TestIsPalindromePermutation(t *testing.T) {
	testInputs := []struct {
		value string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"ab", false},
		{"aba", true},
		{"abab", true},
		{"dabcad", false},
		{"abcaa", false},
		{"Dkd", true},
		{"aaacb", false},
	}
	for _, input := range testInputs {
		got := IsPalindromePermutation(input.value)
		AssertInputs(t, input, got, input.want)
	}
}
