package string

import "strings"

func longestCommonPrefix(strs []string) string {

	var result string

	for index, value := range strs {
		if index == 0 {
			result = value
			continue
		}
		if result == "" {
			break
		}

		for {
			if !strings.HasPrefix(value, result) {
				if len(result) > 0 {
					result = result[:len(result)-1]
				}
			} else {
				break
			}
		}

	}
	return result
}
