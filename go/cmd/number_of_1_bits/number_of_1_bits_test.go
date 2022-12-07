package number_of_1_bits

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNumberOf1Bits_case1(t *testing.T) {
	s := "00000000000000000000000000001011"
	u64, _ := strconv.ParseUint(s, 2, 32)
	testInput := uint32(u64)
	expectedResult := 3

	result := hammingWeight(testInput)
	if result != expectedResult {
		t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", expectedResult, result))
	}
}

func TestNumberOf1Bits_case2(t *testing.T) {
	s := "00000000000000000000000010000000"
	u64, _ := strconv.ParseUint(s, 2, 32)
	testInput := uint32(u64)
	expectedResult := 1

	result := hammingWeight(testInput)
	if result != expectedResult {
		t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", expectedResult, result))
	}
}

func TestNumberOf1Bits_case3(t *testing.T) {
	s := "11111111111111111111111111111101"
	u64, _ := strconv.ParseUint(s, 2, 32)
	testInput := uint32(u64)
	expectedResult := 31

	result := hammingWeight(testInput)
	if result != expectedResult {
		t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", expectedResult, result))
	}
}

func TestNumberOf1Bits_case4(t *testing.T) {
	s := "00000000000000000000000000000001"
	u64, _ := strconv.ParseUint(s, 2, 32)
	testInput := uint32(u64)
	expectedResult := 1

	result := hammingWeight(testInput)
	if result != expectedResult {
		t.Fatal(fmt.Sprintf("Expected `%d`, received `%d`", expectedResult, result))
	}
}
