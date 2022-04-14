package sort

/*

函式名：quicksort
作用：快速排序演算法
引數：
返回值：無
模擬:
   begin:[]int{12, 85, 25, 16, 34, 23, 49, 95, 17, 61}
   -->[],12,[25 16 34 23 49 95 17 61]
   ---->[23 16 17],25,[95 49 61 85]
   ------>[17 16],23,[]
   -------->[16],17,[]
   ---------->[],16,[]
   ------>[34 49 61 85],95,[]
   -------->[],34,[61 85 95]
   ---------->[49],61,[95]
   ------------>[],49,[]
   ------------>[85],95,[]
   -------------->[],85,[]
   last:[]int{12, 16, 17, 23, 25, 34, 49, 61, 85, 95}
*/

func QuickSort(question []int, begin, end int) []int {
	//begin:[]int{12, 85, 25, 16, 34, 23, 49, 95, 17, 61}
	if begin < end {
		key := question[(begin+end)/2]
		i := begin
		j := end

		for {
			for question[i] < key {
				i++
			}
			for question[j] > key {
				j--
			}
			if i >= j {
				break
			}
			question[i], question[j] = question[j], question[i]
		}

		QuickSort(question, begin, i-1)
		QuickSort(question, j+1, end)
	}
	return question
}
