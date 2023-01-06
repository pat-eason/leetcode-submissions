package max_ice_cream_bars

import "sort"

// @see https://leetcode.com/problems/maximum-ice-cream-bars/description/
func maxIceCream(costs []int, coins int) int {
	// sort the array
	sort.Slice(costs, func(i, j int) bool {
		return costs[i] < costs[j]
	})

	// iterate until coins are exhausted
	output := 0
	for i := 0; i < len(costs); i++ {
		// break if coins are exhausted
		if coins < costs[i] {
			break
		}
		coins -= costs[i]
		output++
	}

	return output
}
