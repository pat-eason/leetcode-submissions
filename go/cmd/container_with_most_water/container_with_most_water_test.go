package container_with_most_water

import "testing"

func TestContainerWithMostWater_case1(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	expectedResult := 49
	result := maxArea(input)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, receieved `%d`", expectedResult, result)
	}
}

func TestContainerWithMostWater_case2(t *testing.T) {
	input := []int{1, 1}
	expectedResult := 1
	result := maxArea(input)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, receieved `%d`", expectedResult, result)
	}
}

func TestContainerWithMostWater_case3(t *testing.T) {
	input := []int{2, 3, 4, 5, 18, 17, 6}
	expectedResult := 17
	result := maxArea(input)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, receieved `%d`", expectedResult, result)
	}
}
