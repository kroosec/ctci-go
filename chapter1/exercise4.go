package main

import "strings"

// Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palin­
// drome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation
// is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.

func IsPalindromePermutation(str string) bool {
	occurences := make(map[rune]int)

	str = strings.ToLower(str)
	for _, char := range str {
		occurences[char]++
	}
	// All chars in a string should have an even number of occurences,
	// except one, at most, if string has an odd length.
	oddChars := len(str) % 2
	for _, value := range occurences {
		if value%2 != 0 {
			oddChars--
			if oddChars < 0 {
				return false
			}
		}
	}
	return true
}
