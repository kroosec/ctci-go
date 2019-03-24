package main

import (
	"fmt"
	"strings"
)

// String Compression: Implement a method to perform basic string compression using the counts
// of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the
// "compressed" string would not become smaller than the original string, your method should return
// the original string. You can assume the string has only uppercase and lowercase letters (a - z).

func countCompress(str string) (count int) {
	for i := 0; i < len(str); {
		n := compressFrom(str, i)
		i += n
		count += 2
	}
	return
}

func compressFrom(str string, start int) (size int) {
	j := start + 1
	for ; j < len(str); j++ {
		if str[j] != str[start] {
			break
		}
	}
	return j - start
}
func CompressString(str string) (compressed string) {
	// We count first so that we don't build String Builder if not needed.
	// Also it helps with allocating adequate buffer length, without
	// dynamically growing it each time.
	count := countCompress(str)
	if len(str) <= count {
		return str
	}
	var sb strings.Builder
	sb.Grow(count)
	for i := 0; i < len(str); {
		n := compressFrom(str, i)
		fmt.Fprintf(&sb, "%c%d", str[i], n)
		i += n
	}
	return sb.String()
}

func main() {
}
