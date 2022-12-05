package middle_of_the_linked_list

import (
	"leetcode/internal/types"
)

// @see https://leetcode.com/problems/middle-of-the-linked-list/description/
func middleNode(head *types.ListNode) *types.ListNode {
	slowIndex := head
	fastIndex := head

	for fastIndex != nil && fastIndex.Next != nil {
		fastIndex = fastIndex.Next.Next
		slowIndex = slowIndex.Next
	}

	return slowIndex
}
