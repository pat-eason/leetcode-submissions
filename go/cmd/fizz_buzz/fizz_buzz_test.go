package fizz_buzz

import (
	"fmt"
	"testing"
)

func TestFizzBuss_case1(t *testing.T) {
	expectedResult := []string{"1", "2", "Fizz"}
	testInput := 3
	result := fizzBuzz(testInput)
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i] != result[i] {
			t.Fatal(fmt.Sprintf("Expected `%s`, received `%s`", expectedResult[i], result[i]))
		}
	}
}

func TestFizzBuss_case2(t *testing.T) {
	expectedResult := []string{"1", "2", "Fizz", "4", "Buzz"}
	testInput := 5
	result := fizzBuzz(testInput)
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i] != result[i] {
			t.Fatal(fmt.Sprintf("Expected `%s`, received `%s`", expectedResult[i], result[i]))
		}
	}
}

func TestFizzBuss_case3(t *testing.T) {
	expectedResult := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	testInput := 15
	result := fizzBuzz(testInput)
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i] != result[i] {
			t.Fatal(fmt.Sprintf("Expected `%s`, received `%s`", expectedResult[i], result[i]))
		}
	}
}
