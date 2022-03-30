package leetcode

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {

	newStringArr := []string{}
	for i := 0; i < len(s); i++ {
		switch {
		case 64 < s[i] && s[i] < 91:
			newStringArr = append(newStringArr, strings.ToLower(s[i:i+1]))
		case 96 < s[i] && s[i] < 123:
			newStringArr = append(newStringArr, s[i:i+1])
		case 47 < s[i] && s[i] < 58:
			newStringArr = append(newStringArr, s[i:i+1])
		}
	}
	fmt.Println(newStringArr)
	//if len(newStringArr) == 1 {
	//	return false
	//}

	for i := 0; i < len(newStringArr)/2; i++ {
		if newStringArr[i] != newStringArr[len(newStringArr)-i-1] {
			return false
		}
	}

	return true

}

//這個寫法有料！
func isPalindrome_new(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 判断 c 是否是字符或者数字
func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
