package determine_if_string_halves_are_alike

import "testing"

func TestUniqueOccurrences_case1(t *testing.T) {
	result := halvesAreAlike("book")
	if !result {
		t.Fatal("Expected result to be true")
	}
}

func TestUniqueOccurrences_case2(t *testing.T) {
	result := halvesAreAlike("textbook")
	if result {
		t.Fatal("Expected result to be false")
	}
}

func TestUniqueOccurrences_case3(t *testing.T) {
	result := halvesAreAlike("rracecar")
	if result {
		t.Fatal("Expected result to be false")
	}
}

func TestUniqueOccurrences_case4(t *testing.T) {
	result := halvesAreAlike("ebeilA")
	if !result {
		t.Fatal("Expected result to be true")
	}
}
