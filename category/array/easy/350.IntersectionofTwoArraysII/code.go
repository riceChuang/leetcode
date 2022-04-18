package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	var checkMap = map[int]int{}

	var more, less []int

	if len(nums1) >= len(nums2) {
		more = nums1
		less = nums2
	} else {
		more = nums2
		less = nums1
	}

	for _, value := range less {
		checkMap[value]++
	}
	var result []int
	for _, value := range more {
		if _, ok := checkMap[value]; ok {
			if checkMap[value] != 0 {
				checkMap[value]--
				result = append(result, value)
			}
		}
	}
	return result
}
