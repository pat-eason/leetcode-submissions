package remove_element

import "testing"

func TestRemoveElement_case1(t *testing.T) {
	input := []int{3, 2, 2, 3}
	val := 3
	output := 2
	outputArr := []int{2, 2}

	result := removeElement(input, val)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
	for i := 0; i < output; i++ {
		if !existsInArray(outputArr, input[i]) {
			t.Fatalf("Expected `%d` to exist in output array", input[i])
		}
	}
}

func TestRemoveElement_case2(t *testing.T) {
	input := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	output := 5
	outputArr := []int{0, 1, 4, 0, 3}

	result := removeElement(input, val)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
	for i := 0; i < output; i++ {
		if !existsInArray(outputArr, input[i]) {
			t.Fatalf("Expected `%d` to exist in output array", input[i])
		}
	}
}

func existsInArray(arr []int, value int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return true
		}
	}
	return false
}
