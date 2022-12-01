package valid_parentheses

import (
	"testing"
)

func TestValidParentheses_case1(t *testing.T) {
	result := isValid("()")
	if !result {
		t.Fatal("Expected result to be true")
	}
}

func TestValidParentheses_case2(t *testing.T) {
	result := isValid("()[]{}")
	if !result {
		t.Fatal("Expected result to be true")
	}
}

func TestValidParentheses_case3(t *testing.T) {
	result := isValid("(]")
	if result {
		t.Fatal("Expected result to be false")
	}
}

func TestValidParentheses_case4(t *testing.T) {
	result := isValid("([)]")
	if result {
		t.Fatal("Expected result to be false")
	}
}

func TestValidParentheses_case5(t *testing.T) {
	result := isValid("[")
	if result {
		t.Fatal("Expected result to be false")
	}
}
