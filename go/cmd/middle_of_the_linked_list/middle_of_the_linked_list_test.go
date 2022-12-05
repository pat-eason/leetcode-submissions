package middle_of_the_linked_list

import (
	"leetcode/internal/types"
	"testing"
)

func TestMiddleOfTheLinkedList_case1(t *testing.T) {
	inputValues := []int{1, 2, 3, 4, 5}
	inputList := createLinkedList(inputValues)

	result := middleNode(&inputList)
	if result.Val != 3 {
		t.Fatal("Expected result to be `3`")
	}
}

func TestMiddleOfTheLinkedList_case2(t *testing.T) {
	inputValues := []int{1, 2, 3, 4, 5, 6}
	inputList := createLinkedList(inputValues)

	result := middleNode(&inputList)
	if result.Val != 4 {
		t.Fatal("Expected result to be `4`")
	}
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
