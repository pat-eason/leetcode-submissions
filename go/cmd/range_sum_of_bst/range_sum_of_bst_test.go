package range_sum_of_bst

import (
	"fmt"
	"leetcode/internal/types"
	"testing"
)

func TestRangeSumOfBST_case1(t *testing.T) {
	expectedResult := 32
	testLow := 7
	testHigh := 15

	testTree := types.TreeNode{
		Val: 10,
		Left: &types.TreeNode{
			Val: 5,
			Left: &types.TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &types.TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &types.TreeNode{
			Val:  15,
			Left: nil,
			Right: &types.TreeNode{
				Val:   18,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := rangeSumBST(&testTree, testLow, testHigh)
	if result != expectedResult {
		t.Fatal(fmt.Sprintf("Expected result to be `%d`, got `%d`", expectedResult, result))
	}
}
