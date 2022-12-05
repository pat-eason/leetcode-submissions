package remove_nth_node_from_end_of_list

import "leetcode/internal/types"

func removeNthFromEnd(head *types.ListNode, n int) *types.ListNode {
	values := make([]int, 0)
	var currHead = head

	for currHead != nil {
		values = append(values, currHead.Val)
		currHead = currHead.Next
	}

	culledValues := removeIndex(values, len(values)-n)
	if len(culledValues) == 0 {
		return nil
	}
	output := createLinkedList(culledValues)
	return &output
}

func createLinkedList(input []int) types.ListNode {
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

func removeIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
