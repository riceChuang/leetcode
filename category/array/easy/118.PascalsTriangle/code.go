package leetcode

func generate(numRows int) [][]int {

	result := [][]int{}
	for i := 1; i <= numRows; i++ {
		array := []int{}

		if len(result) == 0 || len(result[len(result)-1]) <= 1 {
			for j := 0; j < i; j++ {
				array = append(array, 1)
			}
			result = append(result, array)
			continue
		}

		var tempStart = 0
		var tempEnd = 1
		for j := 1; j <= i; j++ {
			if j == 1 || j == i {
				array = append(array, 1)
				continue
			}

			value := result[len(result)-1][tempStart] + result[len(result)-1][tempEnd]
			tempStart++
			tempEnd++
			array = append(array, value)
		}
		result = append(result, array)
	}
	return result
}
