package find_first_index_occurence_in_string

import (
	"testing"
)

func TestFoundIndex_case1(t *testing.T) {
	haystack := "sadbutsad"
	needle := "sad"
	expectedIndex := 0

	result := strStr(haystack, needle)

	if result != expectedIndex {
		t.Fatalf("Expected result to equal %d, received %d", expectedIndex, result)
	}
}

func TestFoundIndex_case2(t *testing.T) {
	haystack := "leetcode"
	needle := "leeto"
	expectedIndex := -1

	result := strStr(haystack, needle)

	if result != expectedIndex {
		t.Fatalf("Expected result to equal %d, received %d", expectedIndex, result)
	}
}

func TestFoundIndex_case3(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	expectedIndex := 2

	result := strStr(haystack, needle)

	if result != expectedIndex {
		t.Fatalf("Expected result to equal %d, received %d", expectedIndex, result)
	}
}

func TestFoundIndex_case4(t *testing.T) {
	haystack := "a"
	needle := "a"
	expectedIndex := 0

	result := strStr(haystack, needle)

	if result != expectedIndex {
		t.Fatalf("Expected result to equal %d, received %d", expectedIndex, result)
	}
}
