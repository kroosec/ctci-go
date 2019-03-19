package main

// Check Permutation: Given two strings, write a method to decide if one is a permutation of the
// other.

func IsPermutation(first, second string) bool {
	// Strings of different lengths can't be permutations of each other.
	if len(first) != len(second) {
		return false
	}

	occurences := map[rune]int{}
	for _, c := range first {
		if _, ok := occurences[c]; !ok {
			occurences[c]++
		}
	}
	for _, c := range second {
		occurences[c]--
		if occurences[c] < 0 {
			return false
		}
	}

	return true
}
