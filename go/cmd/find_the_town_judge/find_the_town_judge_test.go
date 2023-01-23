package find_the_town_judge

import "testing"

func TestFindTheTownJudge_case1(t *testing.T) {
	testN := 2
	testTrust := [][]int{{1, 2}}
	expectedResult := 2

	result := findJudge(testN, testTrust)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, received `%d`", expectedResult, result)
	}
}

func TestFindTheTownJudge_case2(t *testing.T) {
	testN := 3
	testTrust := [][]int{{1, 3}, {2, 3}}
	expectedResult := 3

	result := findJudge(testN, testTrust)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, received `%d`", expectedResult, result)
	}
}

func TestFindTheTownJudge_case3(t *testing.T) {
	testN := 3
	testTrust := [][]int{{1, 3}, {2, 3}, {3, 1}}
	expectedResult := -1

	result := findJudge(testN, testTrust)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, received `%d`", expectedResult, result)
	}
}

func TestFindTheTownJudge_case4(t *testing.T) {
	testN := 1
	testTrust := [][]int{}
	expectedResult := 1

	result := findJudge(testN, testTrust)
	if result != expectedResult {
		t.Fatalf("Expected `%d`, received `%d`", expectedResult, result)
	}
}
