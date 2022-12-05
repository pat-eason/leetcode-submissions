package remove_nth_node_from_end_of_list

import (
	"fmt"
	"leetcode/internal/utils"
	"testing"
)

func TestRemoveNthNodeFromEndOfList_case1(t *testing.T) {
	inputValues := []int{1, 2, 3, 4, 5}
	inputN := 2
	inputList := utils.CreateLinkedList(inputValues)
	expectedOutput := []int{1, 2, 3, 5}

	result := removeNthFromEnd(&inputList, inputN)
	curHead := result
	for i := 0; i < len(expectedOutput); i++ {
		if curHead.Val != expectedOutput[i] {
			t.Fatal(fmt.Sprintf("Expected result to be [%d], received [%d]", expectedOutput[i], curHead.Val))
		}
		curHead = curHead.Next
	}
	for curHead != nil {
		t.Fatal("Expected list to be exhausted")
	}
}

func TestRemoveNthNodeFromEndOfList_case2(t *testing.T) {
	inputValues := []int{1}
	inputN := 1
	inputList := utils.CreateLinkedList(inputValues)
	expectedOutput := []int{}

	result := removeNthFromEnd(&inputList, inputN)
	curHead := result
	for i := 0; i < len(expectedOutput); i++ {
		if curHead.Val != expectedOutput[i] {
			t.Fatal(fmt.Sprintf("Expected result to be [%d], received [%d]", expectedOutput[i], curHead.Val))
		}
		curHead = curHead.Next
	}
	for curHead != nil {
		t.Fatal("Expected list to be exhausted")
	}
}

func TestRemoveNthNodeFromEndOfList_case3(t *testing.T) {
	inputValues := []int{1, 2}
	inputN := 1
	inputList := utils.CreateLinkedList(inputValues)
	expectedOutput := []int{1}

	result := removeNthFromEnd(&inputList, inputN)
	curHead := result
	for i := 0; i < len(expectedOutput); i++ {
		if curHead.Val != expectedOutput[i] {
			t.Fatal(fmt.Sprintf("Expected result to be [%d], received [%d]", expectedOutput[i], curHead.Val))
		}
		curHead = curHead.Next
	}
	for curHead != nil {
		t.Fatal("Expected list to be exhausted")
	}
}

func TestRemoveNthNodeFromEndOfList_case4(t *testing.T) {
	inputValues := []int{1, 2}
	inputN := 2
	inputList := utils.CreateLinkedList(inputValues)
	expectedOutput := []int{2}

	result := removeNthFromEnd(&inputList, inputN)
	curHead := result
	for i := 0; i < len(expectedOutput); i++ {
		if curHead.Val != expectedOutput[i] {
			t.Fatal(fmt.Sprintf("Expected result to be [%d], received [%d]", expectedOutput[i], curHead.Val))
		}
		curHead = curHead.Next
	}
	for curHead != nil {
		t.Fatal("Expected list to be exhausted")
	}
}

func TestRemoveNthNodeFromEndOfList_case5(t *testing.T) {
	inputValues := []int{1, 2}
	inputN := 1
	inputList := utils.CreateLinkedList(inputValues)
	expectedOutput := []int{1}

	result := removeNthFromEnd(&inputList, inputN)
	curHead := result
	for i := 0; i < len(expectedOutput); i++ {
		if curHead.Val != expectedOutput[i] {
			t.Fatal(fmt.Sprintf("Expected result to be [%d], received [%d]", expectedOutput[i], curHead.Val))
		}
		curHead = curHead.Next
	}
	for curHead != nil {
		t.Fatal("Expected list to be exhausted")
	}
}
