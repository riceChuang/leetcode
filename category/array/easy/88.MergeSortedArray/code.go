package leetcode

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	for i := 0; i < n; i++ {
		nums1 = append(nums1, nums2[i])
	}
	sort.Ints(nums1)
	fmt.Printf("答案:%v \n", nums1)
}
