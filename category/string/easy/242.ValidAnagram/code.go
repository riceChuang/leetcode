package leetcode

func isAnagram(s string, t string) bool {

	var checkMap = map[string]int{}
	for i := 0; i < len(s); i++ {
		checkMap[s[i:i+1]]++
	}
	for i := 0; i < len(t); i++ {
		checkMap[t[i:i+1]]--
	}
	for _, value := range checkMap {
		if value != 0 {
			return false
		}
	}
	return true
}
