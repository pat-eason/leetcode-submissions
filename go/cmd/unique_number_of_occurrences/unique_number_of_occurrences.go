package unique_number_of_occurrences

import "sort"

func uniqueOccurrences(arr []int) bool {
	sort.Ints(arr)

	occurrences := make(map[int]int)
	for _, i2 := range arr {
		existingOccurrence, present := occurrences[i2]
		value := 0
		if present {
			value = existingOccurrence
		}
		occurrences[i2] = value + 1
	}

	var unique []int
	for _, value := range occurrences {
		if contains(unique, value) {
			return false
		}
		unique = append(unique, value)
	}

	return true
}

func contains(arr []int, lookupValue int) bool {
	// iterate over the array and compare given string to each element
	for _, value := range arr {
		if value == lookupValue {
			return true
		}
	}
	return false
}
