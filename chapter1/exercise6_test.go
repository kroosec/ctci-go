package main

import "testing"

func TestCompressString(t *testing.T) {
	testInputs := []struct {
		str  string
		want string
	}{
		{"", ""},
		{"a", "a"},
		{"ab", "ab"},
		{"aaabb", "a3b2"},
		{"aabcccccaaa", "a2b1c5a3"},
		{"babbbbddd", "b1a1b4d3"},
	}
	for _, input := range testInputs {
		got := CompressString(input.str)
		AssertInputs(t, input, got, input.want)
	}
}
