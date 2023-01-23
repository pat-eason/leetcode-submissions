package find_the_town_judge

// @see https://leetcode.com/problems/find-the-town-judge/
func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n == 1 {
		return n
	}

	whoTrusted := []int{}
	isTrusted := make(map[int]int)

	// iterate over trust
	for i := 0; i < len(trust); i++ {
		if !contains(whoTrusted, trust[i][0]) {
			whoTrusted = append(whoTrusted, trust[i][0])
		}
		isTrusted[trust[i][1]] = isTrusted[trust[i][1]] + 1
	}

	if len(whoTrusted) == n {
		return -1
	}

	for k, v := range isTrusted {
		if v == n-1 && !contains(whoTrusted, k) {
			return k
		}
	}

	return -1
}

func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
