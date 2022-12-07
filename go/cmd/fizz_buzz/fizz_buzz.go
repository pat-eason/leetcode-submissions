package fizz_buzz

import "fmt"

// @see https://leetcode.com/problems/fizz-buzz/description/
func fizzBuzz(n int) []string {
	output := make([]string, 0)
	for i := 1; i <= n; i++ {
		record := ""

		if i%3 == 0 {
			record = "Fizz"
		}

		if i%5 == 0 {
			record = record + "Buzz"
		}

		if len(record) == 0 {
			record = fmt.Sprint(i)
		}

		output = append(output, record)
	}
	return output
}
