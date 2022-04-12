package leetcode

func missingNumber(nums []int) int {

	var result int
	for i := len(nums); i > 0; i-- {
		result += i
		result -= nums[i-1]
	}
	return result
}
