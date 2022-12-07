package range_sum_of_bst

import "leetcode/internal/types"

// @see https://leetcode.com/problems/range-sum-of-bst/description/
func rangeSumBST(root *types.TreeNode, low int, high int) int {
	output := 0

	if root.Val >= low && root.Val <= high {
		output = output + root.Val
	}

	if root.Left != nil {
		output = output + rangeSumBST(root.Left, low, high)
	}
	if root.Right != nil {
		output = output + rangeSumBST(root.Right, low, high)
	}

	return output
}
