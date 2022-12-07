package add_two_numbers

import (
	"leetcode/internal/types"
)

// @see https://leetcode.com/problems/add-two-numbers/description/
func addTwoNumbers(l1 *types.ListNode, l2 *types.ListNode) *types.ListNode {
	output := types.ListNode{
		Val:  0,
		Next: nil,
	}

	carry := 0
	currHead := &output
	l1Head := l1
	l2Head := l2
	for l1Head != nil || l2Head != nil {
		res := 0
		if l1Head != nil {
			res += l1Head.Val
		}
		if l2Head != nil {
			res += l2Head.Val
		}
		if carry > 0 {
			res = res + carry
			carry = 0
		}
		if res >= 10 {
			carry = 1
			res -= 10
		}

		currHead.Val = res
		if (l1Head != nil && l1Head.Next != nil) || (l2Head != nil && l2Head.Next != nil) {
			currHead.Next = &types.ListNode{
				Val:  0,
				Next: nil,
			}
			currHead = currHead.Next
		}

		if l1Head != nil {
			l1Head = l1Head.Next
		}
		if l2Head != nil {
			l2Head = l2Head.Next
		}
	}

	if carry > 0 {
		currHead.Next = &types.ListNode{
			Val:  carry,
			Next: nil,
		}
	}

	return &output
}
