package main

import "testing"

func TestStringRotation(t *testing.T) {
	testTable := []struct {
		str1	string
		str2	string
		want	bool
	}{
		{"", "", true},
		{"a", "", false},
		{"ba", "ab", true},
		{"abc", "abc", true},
		{"abc", "bca", true},
		{"abc", "cab", true},
		{"abc", "cba", false},
		{"abcd", "cdab", true},
		{"abcd", "bcda", true},
		{"abcd", "bcba", false},
		{"abcd", "bcad", false},
	}
	for _, testCase := range testTable {
		got := StringRotation(testCase.str1, testCase.str2)
		AssertInputs(t, testCase, got, testCase.want)
	}
}
