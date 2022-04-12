package leetcode

func firstUniqChar(s string) int {

	var resultMap = map[string][]int{}
	var result = 0
	for i := 0; i < len(s); i++ {
		if _, ok := resultMap[s[i:i+1]]; !ok {
			resultMap[s[i:i+1]] = []int{}
		}
		resultMap[s[i:i+1]] = append(resultMap[s[i:i+1]], i)
	}

	for _, r := range resultMap {

		if len(r) > 1 {
			continue
		}
		if r[0] < result {
			result = r[0]
		}
	}
	return result
}
