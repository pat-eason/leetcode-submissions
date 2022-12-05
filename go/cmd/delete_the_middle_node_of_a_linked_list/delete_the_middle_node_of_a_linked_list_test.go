package delete_the_middle_node_of_a_linked_list

import (
	"fmt"
	"leetcode/internal/utils"
	"testing"
)

func TestDeleteTheMiddleNodeOfALinkedList_case1(t *testing.T) {
	inputValues := []int{1, 3, 4, 7, 1, 2, 6}
	inputList := utils.CreateLinkedList(inputValues)
	expectedOutput := []int{1, 3, 4, 1, 2, 6}

	result := deleteMiddle(&inputList)
	curIndex := 0
	curHead := result
	for curHead != nil {
		if curHead.Val != expectedOutput[curIndex] {
			t.Fatal(fmt.Sprintf("Expected result to be [%d], received [%d]", expectedOutput[curIndex], curHead.Val))
		}
		curIndex++
		curHead = curHead.Next
	}
}

func TestDeleteTheMiddleNodeOfALinkedList_case2(t *testing.T) {
	inputValues := []int{1, 2, 3, 4}
	inputList := utils.CreateLinkedList(inputValues)
	expectedOutput := []int{1, 2, 4}

	result := deleteMiddle(&inputList)
	curIndex := 0
	curHead := result
	for curHead != nil {
		if curHead.Val != expectedOutput[curIndex] {
			t.Fatal(fmt.Sprintf("Expected result to be [%d], received [%d]", expectedOutput[curIndex], curHead.Val))
		}
		curIndex++
		curHead = curHead.Next
	}
}

func TestDeleteTheMiddleNodeOfALinkedList_case3(t *testing.T) {
	inputValues := []int{2, 1}
	inputList := utils.CreateLinkedList(inputValues)
	expectedOutput := []int{2}

	result := deleteMiddle(&inputList)
	curIndex := 0
	curHead := result
	for curHead != nil {
		if curHead.Val != expectedOutput[curIndex] {
			t.Fatal(fmt.Sprintf("Expected result to be [%d], received [%d]", expectedOutput[curIndex], curHead.Val))
		}
		curIndex++
		curHead = curHead.Next
	}
}

func TestDeleteTheMiddleNodeOfALinkedList_case4(t *testing.T) {
	inputValues := []int{1}
	inputList := utils.CreateLinkedList(inputValues)
	expectedOutput := []int{}

	result := deleteMiddle(&inputList)
	curIndex := 0
	curHead := result
	for curHead != nil {
		if curHead.Val != expectedOutput[curIndex] {
			t.Fatal(fmt.Sprintf("Expected result to be [%d], received [%d]", expectedOutput[curIndex], curHead.Val))
		}
		curIndex++
		curHead = curHead.Next
	}
}
