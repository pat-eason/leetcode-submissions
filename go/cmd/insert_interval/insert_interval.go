package insert_interval

import "sort"

// @see https://leetcode.com/problems/insert-interval/description/
func insert(intervals [][]int, newInterval []int) [][]int {
	var output [][]int

	isInserting := false
	insertStart := 0
	insertEnd := 0
	hasInserted := false
	for i := 0; i < len(intervals); i++ {
		isOverlapping := newInterval[0] <= intervals[i][1]
		if newInterval[0] <= intervals[i][0] {
			isOverlapping = newInterval[1] >= intervals[i][0]
		}

		if isOverlapping && !isInserting {
			isInserting = true
			insertStart = newInterval[0]
			if intervals[i][0] < insertStart {
				insertStart = intervals[i][0]
			}
		}

		if isOverlapping && isInserting {
			insertEnd = newInterval[1]
			if intervals[i][1] > insertEnd {
				insertEnd = intervals[i][1]
			}
			continue
		}

		if isInserting {
			isInserting = false
			output = append(output, []int{insertStart, insertEnd})
			hasInserted = true
		}

		output = append(output, intervals[i])
	}

	if isInserting {
		output = append(output, []int{insertStart, insertEnd})
		hasInserted = true
	}

	if len(output) == 0 || !hasInserted {
		output = append(output, newInterval)
	}

	sort.Slice(output[:], func(i, j int) bool {
		for x := range output[i] {
			if output[i][x] == output[j][x] {
				continue
			}
			return output[i][x] < output[j][x]
		}
		return false
	})

	return output
}
