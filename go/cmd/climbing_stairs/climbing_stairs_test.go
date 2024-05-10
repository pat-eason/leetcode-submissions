package climbing_stairs

import "testing"

func TestClimbingStairs_case1(t *testing.T) {
	stairCount := 2
	expectedResult := 2
	result := climbStairs(stairCount)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, received `%d`", expectedResult, result)
	}
}

func TestClimbingStairs_case2(t *testing.T) {
	stairCount := 3
	expectedResult := 3
	result := climbStairs(stairCount)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, received `%d`", expectedResult, result)
	}
}

func TestClimbingStairs_case3(t *testing.T) {
	stairCount := 4
	expectedResult := 5
	result := climbStairs(stairCount)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, received `%d`", expectedResult, result)
	}
}

func TestClimbingStairs_case4(t *testing.T) {
	stairCount := 1
	expectedResult := 1
	result := climbStairs(stairCount)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, received `%d`", expectedResult, result)
	}
}

func TestClimbingStairs_case5(t *testing.T) {
	stairCount := 5
	expectedResult := 8
	result := climbStairs(stairCount)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, received `%d`", expectedResult, result)
	}
}
