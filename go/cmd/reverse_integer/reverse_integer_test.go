package reverse_integer

import "testing"

func TestReverseInteger_case1(t *testing.T) {
	input := 123
	output := 321
	result := reverse(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestReverseInteger_case2(t *testing.T) {
	input := -123
	output := -321
	result := reverse(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestReverseInteger_case3(t *testing.T) {
	input := 120
	output := 21
	result := reverse(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}

func TestReverseInteger_case4(t *testing.T) {
	input := 1534236469
	output := 0
	result := reverse(input)
	if result != output {
		t.Fatalf("Expected `%d`, received `%d`", output, result)
	}
}
