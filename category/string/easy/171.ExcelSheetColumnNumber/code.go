package leetcode

import (
	"math"
)

func titleToNumber(columnTitle string) int {
	runes := []rune(columnTitle)
	var result int
	for i, value := range runes {
		m := math.Pow(26, float64(len(columnTitle)-1-i))
		toInt := (int(value) - 64) * int(m)
		result += toInt
	}
	return result
}
