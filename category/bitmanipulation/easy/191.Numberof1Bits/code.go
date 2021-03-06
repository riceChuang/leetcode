package leetcode

func hammingWeight(num uint32) int {
	var result int
	for num > 0 {
		if num&1 == 1 {
			result++
		}
		num >>= 1
	}
	return result
}
