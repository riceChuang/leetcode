package leetcode

import "math"

func isHappy(n int) bool {
	record := map[int]bool{}
	for n != 1 {
		record[n] = true
		n = getValue(n)
		if _, ok := record[n]; ok {
			return false
		}
	}
	return true
}

func getValue(value int) int {
	var result int
	for value > 0 {
		v := value % 10
		result += int(math.Pow(float64(v), 2))
		value = value / 10
	}
	return result
}
