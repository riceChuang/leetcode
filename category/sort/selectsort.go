package sort

/*
走過整個陣列 每一次選擇當前最小的數值 跟index做交換

*/

func SelectionSort(question []int) []int {

	for i := 0; i < len(question); i++ {
		var minIdx = i
		for j := i; j < len(question); j++ {
			if question[j] < question[minIdx] {
				minIdx = j
			}
		}
		question[i], question[minIdx] = question[minIdx], question[i]
	}

	return question
}
