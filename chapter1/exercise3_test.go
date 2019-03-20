package main

import "testing"

func TestURLify(t *testing.T) {
	inputURLified := []struct {
		str        string
		urlified   string
	}{

		{"", ""},
		{"a", "a"},
		{" ", "%20"},
		{"   ", "%20%20%20"},
		{"a b", "a%20b"},
		{"ab  cd", "ab%20%20cd"},
		{" a b c ", "%20a%20b%20c%20"},
	}

	for _, input := range inputURLified {
		// chars array has enough space.
		numSpaces := CountSpaces([]rune(input.str), len(input.str))
		chars := make([]rune, len(input.str) + numSpaces * 2)
		for i, v := range input.str {
			chars[i] = v
		}
		URLify(chars, len(input.str))
		AssertInputs(t, input.str, string(chars), input.urlified)
	}
}
