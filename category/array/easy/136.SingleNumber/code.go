package leetcode

//ιη¨ XOR δΎε

func singleNumber(nums []int) int {

	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}
