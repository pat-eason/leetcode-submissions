package middle_of_the_linked_list

import (
	"leetcode/internal/utils"
	"testing"
)

func TestMiddleOfTheLinkedList_case1(t *testing.T) {
	inputValues := []int{1, 2, 3, 4, 5}
	inputList := utils.CreateLinkedList(inputValues)

	result := middleNode(&inputList)
	if result.Val != 3 {
		t.Fatal("Expected result to be `3`")
	}
}

func TestMiddleOfTheLinkedList_case2(t *testing.T) {
	inputValues := []int{1, 2, 3, 4, 5, 6}
	inputList := utils.CreateLinkedList(inputValues)

	result := middleNode(&inputList)
	if result.Val != 4 {
		t.Fatal("Expected result to be `4`")
	}
}
