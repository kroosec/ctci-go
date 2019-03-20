package main

// URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string
// has sufficient space at the end to hold the additional characters, and that you are given the "true"
// length of the string. (Note: If implementing in Java, please use a character array so that you can
// perform this operation in place.)
// EXAMPLE
// Input: "Mr John Smith	", 13
// Output: "Mr%20John%20Smith"

func CountSpaces(str []rune, initialLen int) (numSpaces int) {
	for i := 0; i < initialLen; i++ {
		if str[i] == rune(' ') {
			numSpaces++
		}
	}
	return
}

// Strings are immutable in Go. We use a slice of runes, to change string in place.
func URLify(str []rune, initialLen int) {
	numSpaces := CountSpaces(str, initialLen)
	j := initialLen + numSpaces*2 - 1
	for i := initialLen - 1; i >= 0; i-- {
		if str[i] == rune(' ') {
			str[j] = '0'
			str[j-1] = '2'
			str[j-2] = '%'
			j = j - 3
		} else {
			str[j] = str[i]
			j--
		}
	}
}
