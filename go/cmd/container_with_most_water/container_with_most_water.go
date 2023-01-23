package container_with_most_water

import "math"

// @see https://leetcode.com/problems/container-with-most-water/description/
func maxArea(height []int) int {
	startIndex := 0
	endIndex := len(height) - 1
	maxResult := 0
	for startIndex < endIndex {
		maxHeight := math.Min(float64(height[startIndex]), float64(height[endIndex]))
		currentArea := int(maxHeight) * (endIndex - startIndex)
		if currentArea > maxResult {
			maxResult = currentArea
		}
		if height[startIndex] > height[endIndex] {
			endIndex--
		} else {
			startIndex++
		}
	}
	return maxResult
}
