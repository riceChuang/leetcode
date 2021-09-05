package twosum1

func twoSum(nums []int, target int) []int {

	for i, a := range nums[:len(nums)-1] {
		for j, b := range nums[i+1:] {
			if a + b == target {
				return []int{i, i+1+j}
			}
		}
	}
	return []int{}
}
