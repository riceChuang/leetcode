package leetcode

import "strconv"

func Solution(cars []string) []int {
	// write your code in Go 1.4
	result := []int{}

	for i, carA := range cars {
		var r int
		for j, carB := range cars {
			if i == j {
				continue
			}
			car1 := Conv(carA)
			car2 := Conv(carB)
			var carResult int
			for car1 != 0 || car2 != 0 {
				res1 := car1 & 1
				res2 := car2 & 1
				car1 >>= 1
				car2 >>= 1
				if res1 != res2 {
					carResult++
				}
			}
			if carResult <= 1 {
				r++
			}
		}
		result = append(result, r)
	}
	return result
}

func Conv(str string) int {
	if i, err := strconv.ParseInt(str, 2, 64); err != nil {
		return 0
	} else {
		return int(i)
	}
}
