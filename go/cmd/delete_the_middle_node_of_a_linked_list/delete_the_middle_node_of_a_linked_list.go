package delete_the_middle_node_of_a_linked_list

import "leetcode/internal/types"

// @see https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list/description/
func deleteMiddle(head *types.ListNode) *types.ListNode {
	var prevHead *types.ListNode = nil
	var slowIndex = head
	var fastIndex = head

	for fastIndex != nil && fastIndex.Next != nil {
		prevHead = slowIndex
		fastIndex = fastIndex.Next.Next
		slowIndex = slowIndex.Next
	}

	if prevHead == nil {
		return nil
	}

	prevHead.Next = slowIndex.Next

	return head
}
