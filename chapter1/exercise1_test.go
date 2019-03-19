package main

import "testing"

var inputIsUnique []struct {
	want bool
	str  string
} = []struct {
	want bool
	str  string
}{
	{true, ""},
	{true, "abcdef"},
	{false, "foo"},
	{false, "barb"},
}

func TestIsUnique(t *testing.T) {
	for _, input := range inputIsUnique {
		AssertBool(t, input.str, IsUnique(input.str), input.want)
		AssertBool(t, input.str, IsUnique(input.str), input.want)
	}
}

func TestIsUniqueNoExtras(t *testing.T) {
	for _, input := range inputIsUnique {
		AssertBool(t, input.str, IsUniqueNoExtras(input.str), input.want)
		AssertBool(t, input.str, IsUniqueNoExtras(input.str), input.want)
	}
}
