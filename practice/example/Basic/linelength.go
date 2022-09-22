package Basic

import "unicode/utf8"

func GetLineLength(words []string) int {
	total := 0
	for _, word := range words {
		total += utf8.RuneCountInString(word)
	}
	numSpaces := len(words) - 1
	return total + numSpaces
}
