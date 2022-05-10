package search

//地回方式
func BinarySearchRecursively(question []int, target int, start int, end int) int {
	if end < start {
		return -1
	}
	midIndex := (start + end) / 2
	middleValue := question[midIndex]

	if target > middleValue {
		return BinarySearchRecursively(question, target, midIndex+1, end)
	} else if target < middleValue {
		return BinarySearchRecursively(question, target, start, midIndex-1)
	} else {
		return middleValue
	}

}

func BinarySearch(question []int, target int) int {
	start := 0
	end := len(question) - 1

	for end >= start {
		mid := (start + end) / 2
		if target == question[mid] {
			return mid
		} else if target > question[mid] {
			start = mid + 1
		} else if target < question[mid] {
			end = mid - 1
		}
	}
	return -1
}
