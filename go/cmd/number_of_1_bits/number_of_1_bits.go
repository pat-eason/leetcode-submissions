package number_of_1_bits

import "strconv"

// @see https://leetcode.com/problems/number-of-1-bits/description/
func hammingWeight(num uint32) int {
	output := 0
	binaryString := strconv.FormatInt(int64(num), 2)

	endIndex := len(binaryString) - 1
	for i := 0; i <= endIndex; i++ {
		output = output + (int(binaryString[i]) - 48)
		if endIndex > i {
			output = output + (int(binaryString[endIndex]) - 48)
		}
		endIndex--
	}
	return output
}
