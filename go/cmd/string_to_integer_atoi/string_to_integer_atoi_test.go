package string_to_integer_atoi

import (
	"fmt"
	"math"
	"testing"
)

func TestInt32Overflow(t *testing.T) {
	test1 := 10 + math.MaxInt32
	test2 := math.MaxInt32
	test3 := -10 - math.MaxInt32
	fmt.Printf("%d | %d | %d", test1, test2, test3)
	fmt.Printf("%v", []rune("0123456789+-"))
	fmt.Printf("%v", []rune("00++00"))
	fmt.Printf("%v", []rune("00++00"))
}

func TestStringToIntegerAtoI_case1(t *testing.T) {
	input := "42"
	output := 42
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case2(t *testing.T) {
	input := "  -42"
	output := -42
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case3(t *testing.T) {
	input := "4193 with words"
	output := 4193
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case4(t *testing.T) {
	input := "123456789123456789123456789"
	output := math.MaxInt32
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case5(t *testing.T) {
	input := "42-456789"
	output := 42456789
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case6(t *testing.T) {
	input := "words and 987"
	output := 0
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case7(t *testing.T) {
	input := "-91283472332"
	output := -2147483648
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case8(t *testing.T) {
	input := "+1"
	output := 1
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case9(t *testing.T) {
	input := "+-12"
	output := 0
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case10(t *testing.T) {
	input := "21474836460"
	output := 2147483647
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case11(t *testing.T) {
	input := "00000-42a1234"
	output := 0
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case12(t *testing.T) {
	input := "  -0012a42"
	output := -12
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case13(t *testing.T) {
	input := "   +0 123"
	output := 0
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case14(t *testing.T) {
	input := "-13+8"
	output := -13
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestStringToIntegerAtoI_case15(t *testing.T) {
	input := "21474836++"
	output := 21474836
	result := myAtoi(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}
