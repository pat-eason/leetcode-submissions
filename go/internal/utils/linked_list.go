package utils

import "leetcode/internal/types"

func CreateLinkedList(input []int) types.ListNode {
	output := types.ListNode{
		Val: input[0],
	}
	curHead := &output
	for i := 1; i < len(input); i++ {
		next := types.ListNode{Val: input[i]}
		curHead.Next = &next
		curHead = &next
	}
	return output
}
