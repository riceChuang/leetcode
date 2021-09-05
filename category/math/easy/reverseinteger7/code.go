package reverseinteger7

import "math"

func reverse(x int) int {

	var result int
	var isNegative bool
	newX := x
	if x > 0 {
		isNegative = true
	} else {
		newX = -x
	}

	for newX > 0 {
		number := newX % 10
		result = result * 10
		result += number
		newX /= 10
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	if !isNegative {
		return -result
	}

	return result
}
