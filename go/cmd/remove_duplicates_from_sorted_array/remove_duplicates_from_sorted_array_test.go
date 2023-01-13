package remove_duplicates_from_sorted_array

import (
	"testing"
)

func TestRemoveDuplicatesFromSortedArray_case1(t *testing.T) {
	input := []int{1, 1, 2}
	output := 2
	outputArr := []int{1, 2}

	result := removeDuplicates(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
	for i := 0; i < len(outputArr); i++ {
		if input[i] != outputArr[i] {
			t.Fatalf("Expected index `%d` to be `%d`, received `%d`", i, outputArr[i], input[i])
		}
	}
}

func TestRemoveDuplicatesFromSortedArray_case2(t *testing.T) {
	input := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	output := 5
	outputArr := []int{0, 1, 2, 3, 4}

	result := removeDuplicates(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
	for i := 0; i < len(outputArr); i++ {
		if input[i] != outputArr[i] {
			t.Fatalf("Expected index `%d` to be `%d`, received `%d`", i, outputArr[i], input[i])
		}
	}
}
