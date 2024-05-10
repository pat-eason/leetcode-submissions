package find_first_index_occurence_in_string

import (
	"strings"
)

// @see https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
func strStr(haystack string, needle string) int {
	haystackRunes := []rune(haystack)

	var haystackStringBuilder strings.Builder
	for i := 0; i < len(haystackRunes); i++ {
		currentRune := haystackRunes[i]
		haystackStringBuilder.WriteRune(currentRune)

		if strings.HasPrefix(needle, haystackStringBuilder.String()) {
			if haystackStringBuilder.String() == needle {
				return i
			}

			for j := i + 1; j < len(haystackRunes); j++ {
				currentRune := haystackRunes[j]
				haystackStringBuilder.WriteRune(currentRune)

				if haystackStringBuilder.String() == needle {
					return i
				}

				if !strings.HasPrefix(needle, haystackStringBuilder.String()) || haystackStringBuilder.Len() == len(needle) {
					break
				}
			}
		}

		haystackStringBuilder.Reset()
	}

	return -1
}
