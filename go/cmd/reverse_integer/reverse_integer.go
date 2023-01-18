package reverse_integer

import (
	"fmt"
	"math"
)

var (
	runeMap = map[rune]int{
		48: 0,
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
	max32     = int(math.Pow(2, 31) - 1)
	min32     = int(math.Pow(-2, 31))
)

// @see https://leetcode.com/problems/reverse-integer/description/
func reverse(x int) int {
	output := 0
	inputRunes := []rune(fmt.Sprint(x))

	isNegative := inputRunes[0] == '-'

	for i := len(inputRunes) - 1; i >= 0; i-- {
		if inputRunes[i] == minusRune {
			continue
		}
		output = output*10 + runeMap[inputRunes[i]]
		if output > max32 || output < min32 {
			return 0
		}
	}

	if isNegative {
		return -output
	}
	return output
}
