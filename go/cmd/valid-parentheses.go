package cmd

import (
	"container/list"
	"fmt"
)

// @see https://leetcode.com/problems/valid-parentheses/description/
func isValid(s string) bool {
	openChars := []string{"(", "[", "{"}
	closeChars := []string{")", "]", "}"}
	openParens := list.New()

	for _, c := range s {
		currentChar := string(c)

		isOpenChar := contains(openChars, currentChar)
		if isOpenChar {
			openParens.PushFront(currentChar)
		} else {
			previousChar := openParens.Front()
			if previousChar == nil {
				return false
			}
			previousCharVal := fmt.Sprintf("%s", previousChar.Value)
			openIndex := getIndex(openChars, previousCharVal)

			if openIndex == -1 || closeChars[openIndex] != currentChar {
				return false
			}

			openParens.Remove(previousChar)
		}
	}

	return openParens.Len() == 0
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func getIndex(s []string, str string) int {
	for i, v := range s {
		if v == str {
			return i
		}
	}
	return -1
}
