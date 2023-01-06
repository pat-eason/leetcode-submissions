package max_ice_cream_bars

import (
	"fmt"
	"testing"
)

func TestMaxIceCreamBars_case1(t *testing.T) {
	testCosts := []int{1, 3, 2, 4, 1}
	testCoins := 7
	expectedResult := 4
	result := maxIceCream(testCosts, testCoins)
	if result != expectedResult {
		t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", expectedResult, result))
	}
}

func TestMaxIceCreamBars_case2(t *testing.T) {
	testCosts := []int{10, 6, 8, 7, 7, 8}
	testCoins := 5
	expectedResult := 0
	result := maxIceCream(testCosts, testCoins)
	if result != expectedResult {
		t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", expectedResult, result))
	}
}

func TestMaxIceCreamBars_case3(t *testing.T) {
	testCosts := []int{1, 6, 3, 1, 2, 5}
	testCoins := 20
	expectedResult := 6
	result := maxIceCream(testCosts, testCoins)
	if result != expectedResult {
		t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", expectedResult, result))
	}
}

func TestMaxIceCreamBars_case4(t *testing.T) {
	testCosts := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	testCoins := 0
	expectedResult := 0
	result := maxIceCream(testCosts, testCoins)
	if result != expectedResult {
		t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", expectedResult, result))
	}
}

func TestMaxIceCreamBars_case5(t *testing.T) {
	testCosts := []int{1, 1}
	testCoins := 50
	expectedResult := 2
	result := maxIceCream(testCosts, testCoins)
	if result != expectedResult {
		t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", expectedResult, result))
	}
}

func TestMaxIceCreamBars_case6(t *testing.T) {
	testCosts := []int{2}
	testCoins := 1
	expectedResult := 0
	result := maxIceCream(testCosts, testCoins)
	if result != expectedResult {
		t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", expectedResult, result))
	}
}
