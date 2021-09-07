package leetcode

func searchInsert(nums []int, target int) int {
	var resultIndex int
	for index, value := range nums {
		resultIndex = index
		if target <= value {
			return resultIndex
		}
	}
	return resultIndex + 1
}
