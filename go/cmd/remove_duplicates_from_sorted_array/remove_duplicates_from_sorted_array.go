package remove_duplicates_from_sorted_array

// @see https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	updateIndex := 0
	for i := 0; i < len(nums); i++ {
		// if next number had not changed then continue
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			continue
		}

		// if number has changed then:
		// 	* update the update index value
		// 	* update the new current number
		// 	* increment the update index
		nums[updateIndex] = nums[i]
		updateIndex++
	}
	return updateIndex
}
