package unique_number_of_occurrences

import "testing"

func TestUniqueOccurrences_case1(t *testing.T) {
	result := uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})
	if !result {
		t.Fatal("Expected result to be true")
	}
}

func TestUniqueOccurrences_case2(t *testing.T) {
	result := uniqueOccurrences([]int{1, 2})
	if result {
		t.Fatal("Expected result to be false")
	}
}

func TestUniqueOccurrences_case3(t *testing.T) {
	result := uniqueOccurrences([]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0})
	if !result {
		t.Fatal("Expected result to be true")
	}
}

func TestUniqueOccurrences_case4(t *testing.T) {
	result := uniqueOccurrences([]int{26, 2, 16, 16, 5, 5, 26, 2, 5, 20, 20, 5, 2, 20, 2, 2, 20, 2, 16, 20, 16, 17, 16, 2, 16, 20, 26, 16})
	if result {
		t.Fatal("Expected result to be false")
	}
}
