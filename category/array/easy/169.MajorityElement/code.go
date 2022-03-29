package leetcode

func majorityElement(nums []int) int {
	var result = make(map[int]int)

	for _, value := range nums {
		result[value]++
		if result[value] > len(nums)/2 {
			return value
		}
	}
	return 0
}

/*
摩爾投票算法 因為題目是 超過50%的數
所以可以考慮一件事情
今天有個數再100票裡面站了
60% (有60個1)
40% (是其他數目)
利用互相抵銷的做法
就會留下 20% 的 1 也就是答案
同
51% (有51個1)
49% (是其他數目)
就會留下 2% 的 1 也就是答案
*/

func majorityElement1(nums []int) int {

	var result, count = nums[0], 0
	for _, value := range nums {

		if count == 0 {
			result, count = value, 1
		} else if result == value {
			count++
		} else {
			count--
		}

	}
	return result
}
