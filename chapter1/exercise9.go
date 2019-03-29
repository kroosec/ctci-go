package main

import "strings"

// String Rotation: Assume you have a method isSubstringwhich checks if one
// word is a substring of another. Given two strings, sl and s2, write code to
// check if s2 is a rotation of sl using only one call to isSubstring (e.g.,
// "waterbottle" is a rotation of"erbottlewat").

// Concatenate one of the strings with itself. The other string would be its substring.
func StringRotation(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	concat := str2 + str2
	return strings.Contains(concat, str1)
}
