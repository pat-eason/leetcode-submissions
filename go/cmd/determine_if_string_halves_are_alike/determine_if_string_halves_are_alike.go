package determine_if_string_halves_are_alike

import (
	"regexp"
)

var vowelRegex = regexp.MustCompile(`(?m)[aeiouAEIOU]`)

func halvesAreAlike(s string) bool {
	leftSideVowels := 0
	rightSideVowels := 0

	rIdx := len(s) - 1
	for lIdx := 0; lIdx < rIdx; lIdx++ {
		if isVowel(string(s[lIdx])) {
			leftSideVowels++
		}

		if isVowel(string(s[rIdx])) {
			rightSideVowels++
		}
		rIdx--
	}

	return leftSideVowels == rightSideVowels
}

func isVowel(input string) bool {
	result := vowelRegex.FindAllString(input, -1)
	return len(result) > 0
}
