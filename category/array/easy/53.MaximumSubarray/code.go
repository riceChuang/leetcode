package leetcode

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum, midSum := nums[0], 0
	for _, value := range nums {
		midSum += value

		if midSum > maxSum {
			maxSum = midSum
		}

		if midSum < 0 {
			midSum = 0
		}
	}
	return maxSum
}
