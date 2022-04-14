package sort

//question [22 15 11 45 13]
//ans [11 13 15 22 45]

func BubbleSort(question []int) []int {
	for i := 0; i < len(question)-1; i++ {
		for j := 0; j < len(question)-i-1; j++ {
			if question[j] > question[j+1] {
				question[j], question[j+1] = question[j+1], question[j]
			}
		}
	}
	return question
}
