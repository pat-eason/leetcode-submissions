package add_two_numbers

import (
	"fmt"
	"leetcode/internal/utils"
	"testing"
)

func TestAddTwoNumbers_case1(t *testing.T) {
	testListA := utils.CreateLinkedList([]int{2, 4, 3})
	testListB := utils.CreateLinkedList([]int{5, 6, 4})
	expectedList := utils.CreateLinkedList([]int{7, 0, 8})

	result := addTwoNumbers(&testListA, &testListB)

	exHead := &expectedList
	resultHead := result

	for exHead != nil {
		if exHead.Val != resultHead.Val {
			t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", exHead.Val, resultHead.Val))
		}
		exHead = exHead.Next
		resultHead = resultHead.Next
	}
}

func TestAddTwoNumbers_case2(t *testing.T) {
	testListA := utils.CreateLinkedList([]int{0})
	testListB := utils.CreateLinkedList([]int{0})
	expectedList := utils.CreateLinkedList([]int{0})

	result := addTwoNumbers(&testListA, &testListB)

	exHead := &expectedList
	resultHead := result

	for exHead != nil {
		if exHead.Val != resultHead.Val {
			t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", exHead.Val, resultHead.Val))
		}
		exHead = exHead.Next
		resultHead = resultHead.Next
	}
}

func TestAddTwoNumbers_case3(t *testing.T) {
	testListA := utils.CreateLinkedList([]int{9, 9, 9, 9, 9, 9, 9})
	testListB := utils.CreateLinkedList([]int{9, 9, 9, 9})
	expectedList := utils.CreateLinkedList([]int{8, 9, 9, 9, 0, 0, 0, 1})

	result := addTwoNumbers(&testListA, &testListB)

	exHead := &expectedList
	resultHead := result

	for exHead != nil {
		if exHead.Val != resultHead.Val {
			t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", exHead.Val, resultHead.Val))
		}
		exHead = exHead.Next
		resultHead = resultHead.Next
	}
}

func TestAddTwoNumbers_case4(t *testing.T) {
	testListA := utils.CreateLinkedList([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	testListB := utils.CreateLinkedList([]int{5, 6, 4})
	expectedList := utils.CreateLinkedList([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})

	result := addTwoNumbers(&testListA, &testListB)

	exHead := &expectedList
	resultHead := result

	for exHead != nil {
		if exHead.Val != resultHead.Val {
			t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", exHead.Val, resultHead.Val))
		}
		exHead = exHead.Next
		resultHead = resultHead.Next
	}
}
