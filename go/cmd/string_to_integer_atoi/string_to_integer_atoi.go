package string_to_integer_atoi

import (
	"math"
)

var (
	charRunes = []rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 43, 45}
	runeMap   = map[rune]int{
		49: 1,
		50: 2,
		51: 3,
		52: 4,
		53: 5,
		54: 6,
		55: 7,
		56: 8,
		57: 9,
	}
	minusRune = int32(45)
	plusRune  = int32(43)
	zeroRune  = int32(48)
	max32     = int(math.Pow(2, 31) - 1)
	min32     = int(math.Pow(-2, 31))
)

// @see https://leetcode.com/problems/string-to-integer-atoi/description/
func myAtoi(s string) int {
	output := 0
	isNegative := false
	nextCharShouldBeNumber := false

	for _, iRune := range s {
		if !runeIsAccepted(iRune) {
			if iRune == ' ' && nextCharShouldBeNumber {
				return formatOutput(output, isNegative)
			}
			if iRune != ' ' || output > 0 {
				return formatOutput(output, isNegative)
			}
			continue
		}

		if (iRune == minusRune || iRune == plusRune) && nextCharShouldBeNumber {
			return formatOutput(output, isNegative)
		}
		nextCharShouldBeNumber = nextCharShouldBeNumber || (iRune == minusRune || iRune == plusRune || iRune == zeroRune)

		if iRune == minusRune && output > 0 {
			continue
		}

		if iRune == minusRune {
			if output == 0 {
				isNegative = true
			}
			continue
		}

		if iRune != minusRune && iRune != plusRune {
			output = output*10 + (runeMap[iRune])
			if output > max32 || output < min32 {
				if isNegative {
					return min32
				}
				return max32
			}
		}
	}

	return formatOutput(output, isNegative)
}

func formatOutput(value int, isNegative bool) int {
	if isNegative {
		return -value
	}
	return value
}

func runeIsAccepted(str rune) bool {
	for _, v := range charRunes {
		if v == str {
			return true
		}
	}
	return false
}
