package main

// Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
// cannot use additional data structures?

// Using a hashtable.
func IsUnique(str string) bool {
	// Assuming ASCII characters, strings longer than 256 can't be unique.
	if len(str) > 256 {
		return false
	}

	exists := make(map[rune]struct{}, len(str))
	for _, s := range str {
		if _, ok := exists[s]; ok {
			return false
		} else {
			exists[s] = struct{}{}
		}
	}
	return true
}

// Without using extra data structures.
func IsUniqueNoExtras(str string) bool {
	if len(str) > 256 {
		return false
	}
	for i := 0; i < len(str); i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}
