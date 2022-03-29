package leetcode

func convertToTitle(columnNumber int) string {
	result := []byte{}

	for columnNumber > 0 {
		result = append(result, 'A'+byte((columnNumber-1)%26))
		columnNumber = (columnNumber - 1) / 26
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return string(result)
}
