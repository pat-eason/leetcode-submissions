package climbing_stairs

// @see https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
	if n == 0 {
		return 0
	}

	remainder1 := n - 1
	remainder2 := n - 1

	if remainder1 == 0 {
		remainder1 = 1
	}
	if remainder2 == 0 {
		remainder2 = 1
	}

	output := remainder1
	if remainder2 > 0 {
		output = output + remainder2
	}
	return output
}
