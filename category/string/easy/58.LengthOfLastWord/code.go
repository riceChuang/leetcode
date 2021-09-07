package leetcode

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	var charLastIndex int
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " && charLastIndex == 0 {
			charLastIndex = i
		}
		if string(s[i]) == " " && charLastIndex != 0 {
			return charLastIndex - i
		}
		if i == 0 {
			return charLastIndex + 1
		}
	}
	return charLastIndex
}
