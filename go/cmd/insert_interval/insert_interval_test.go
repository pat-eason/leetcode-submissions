package insert_interval

import (
	"fmt"
	"testing"
)

func TestInsertInterval_case1(t *testing.T) {
	inputIntervals := [][]int{{1, 3}, {6, 9}}
	inputNewInterval := []int{2, 5}
	expectedResult := [][]int{{1, 5}, {6, 9}}

	result := insert(inputIntervals, inputNewInterval)
	if len(result) != len(expectedResult) {
		t.Fatalf("Resulting array is not correct length. Expected `%d`, received `%d`", len(expectedResult), len(result))
	}
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i][0] != result[i][0] || expectedResult[i][1] != result[i][1] {
			t.Fatalf("Expected `%d`, received `%d` on index `%d`", expectedResult[i], result[i], i)
		}
	}
}

func TestInsertInterval_case2(t *testing.T) {
	inputIntervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	inputNewInterval := []int{4, 8}
	expectedResult := [][]int{{1, 2}, {3, 10}, {12, 16}}

	result := insert(inputIntervals, inputNewInterval)
	if len(result) != len(expectedResult) {
		t.Fatalf("Resulting array is not correct length. Expected `%d`, received `%d`", len(expectedResult), len(result))
	}
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i][0] != result[i][0] || expectedResult[i][1] != result[i][1] {
			t.Fatalf("Expected `%d`, received `%d` on index `%d`", expectedResult[i], result[i], i)
		}
	}
}

func TestInsertInterval_case3(t *testing.T) {
	inputIntervals := [][]int{}
	inputNewInterval := []int{5, 7}
	expectedResult := [][]int{{5, 7}}

	result := insert(inputIntervals, inputNewInterval)
	if len(result) != len(expectedResult) {
		t.Fatalf("Resulting array is not correct length. Expected `%d`, received `%d`", len(expectedResult), len(result))
	}
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i][0] != result[i][0] || expectedResult[i][1] != result[i][1] {
			t.Fatalf("Expected `%d`, received `%d` on index `%d`", expectedResult[i], result[i], i)
		}
	}
}

func TestInsertInterval_case4(t *testing.T) {
	inputIntervals := [][]int{{1, 5}}
	inputNewInterval := []int{2, 3}
	expectedResult := [][]int{{1, 5}}

	result := insert(inputIntervals, inputNewInterval)
	if len(result) != len(expectedResult) {
		t.Fatalf("Resulting array is not correct length. Expected `%d`, received `%d`", len(expectedResult), len(result))
	}
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i][0] != result[i][0] || expectedResult[i][1] != result[i][1] {
			t.Fatalf("Expected `%d`, received `%d` on index `%d`", expectedResult[i], result[i], i)
		}
	}
}

func TestInsertInterval_case5(t *testing.T) {
	inputIntervals := [][]int{{1, 5}}
	inputNewInterval := []int{5, 7}
	expectedResult := [][]int{{1, 7}}

	result := insert(inputIntervals, inputNewInterval)
	if len(result) != len(expectedResult) {
		t.Fatalf("Resulting array is not correct length. Expected `%d`, received `%d`", len(expectedResult), len(result))
	}
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i][0] != result[i][0] || expectedResult[i][1] != result[i][1] {
			t.Fatal(fmt.Sprintf("Expected `%d`, received `%d` on index `%d`", expectedResult[i], result[i], i))
		}
	}
}

func TestInsertInterval_case6(t *testing.T) {
	inputIntervals := [][]int{{1, 5}}
	inputNewInterval := []int{6, 8}
	expectedResult := [][]int{{1, 5}, {6, 8}}

	result := insert(inputIntervals, inputNewInterval)
	if len(result) != len(expectedResult) {
		t.Fatalf("Resulting array is not correct length. Expected `%d`, received `%d`", len(expectedResult), len(result))
	}
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i][0] != result[i][0] || expectedResult[i][1] != result[i][1] {
			t.Fatalf("Expected `%d`, received `%d` on index `%d`", expectedResult[i], result[i], i)
		}
	}
}

func TestInsertInterval_case7(t *testing.T) {
	inputIntervals := [][]int{{1, 5}}
	inputNewInterval := []int{0, 3}
	expectedResult := [][]int{{0, 5}}

	result := insert(inputIntervals, inputNewInterval)
	if len(result) != len(expectedResult) {
		t.Fatalf("Resulting array is not correct length. Expected `%d`, received `%d`", len(expectedResult), len(result))
	}
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i][0] != result[i][0] || expectedResult[i][1] != result[i][1] {
			t.Fatalf("Expected `%d`, received `%d` on index `%d`", expectedResult[i], result[i], i)
		}
	}
}

func TestInsertInterval_case8(t *testing.T) {
	inputIntervals := [][]int{{1, 5}}
	inputNewInterval := []int{0, 0}
	expectedResult := [][]int{{0, 0}, {1, 5}}

	result := insert(inputIntervals, inputNewInterval)
	if len(result) != len(expectedResult) {
		t.Fatalf("Resulting array is not correct length. Expected `%d`, received `%d`", len(expectedResult), len(result))
	}
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i][0] != result[i][0] || expectedResult[i][1] != result[i][1] {
			t.Fatalf("Expected `%d`, received `%d` on index `%d`", expectedResult[i], result[i], i)
		}
	}
}

func TestInsertInterval_case9(t *testing.T) {
	inputIntervals := [][]int{{3, 5}, {12, 15}}
	inputNewInterval := []int{6, 6}
	expectedResult := [][]int{{3, 5}, {6, 6}, {12, 15}}

	result := insert(inputIntervals, inputNewInterval)
	if len(result) != len(expectedResult) {
		t.Fatalf("Resulting array is not correct length. Expected `%d`, received `%d`", len(expectedResult), len(result))
	}
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i][0] != result[i][0] || expectedResult[i][1] != result[i][1] {
			t.Fatalf("Expected `%d`, received `%d` on index `%d`", expectedResult[i], result[i], i)
		}
	}
}
