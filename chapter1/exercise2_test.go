package main

import "testing"

var inputIsPermutation []struct {
	want bool
	first  string
	second  string
} = []struct {
	want bool
	first  string
	second  string
}{
	{true, "", ""},
	{true, "a", "a"},
	{true, "ab", "ba"},
	{true, "cab", "bca"},
	{true, "b c", "cb "},
	{false, "a", "b"},
	{false, "b", ""},
	{false, "ab", "ac"},
	{false, "cab", "bcao"},
	{false, "cabc", "bcao"},
	{false, "cabc", "bcab"},
	{false, "abcdefg", "gfdbcaa"},
}

func TestIsPermutation(t *testing.T) {
	for _, input := range inputIsPermutation {
		got := IsPermutation(input.first, input.second)
		AssertBool(t, input, got, input.want)
	}
}
