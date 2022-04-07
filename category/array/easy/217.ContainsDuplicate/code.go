package leetcode

func containsDuplicate(nums []int) bool {

	record := map[int]int{}

	for _, value := range nums {
		record[value]++
		if record[value] >= 2 {
			return true
		}
	}
	return false
}