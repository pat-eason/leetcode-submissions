package remove_element

func removeElement(nums []int, val int) int {
	updateIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[updateIndex] = nums[i]
		updateIndex++
	}
	return updateIndex
}
