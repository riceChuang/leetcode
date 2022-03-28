package leetcode

import "fmt"

/*
動態規劃(Dynamic Programming)
	假定要做一個 5!的計算
	5*4*3*2*1 = 120
*/

//正常用迴圈的解法
const N = 5

func factorial() {
	result := 1
	for i := 2; i <= N; i++ {
		result *= i
	}
	fmt.Println(result)
}

//用動態規劃的方式
func factorial_DP() {
	var result = make([]int, 10)
	result[0] = 0
	result[1] = 1
	for i := 2; i <= N; i++ {
		result[i] = result[i-1] * i
	}
	fmt.Println(result)
}

/*
	費布那西數列（Fibonacci)
	=> 0
	=> 1
	=> 0+1 = 1
	=> 1+1 = 2
	=> 1+2 = 3
	=> 2+3 = 5
	=> 3+5 = 8
	....
*/

//時間複雜度為 O(N)
func fibonacci_DP(target int) int {
	var result = make([]int, target+1)
	result[0] = 1
	result[1] = 1
	for i := 2; i <= target; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[target-1]
}

//時間複雜度為 O(2 n次方)
func fibonacci_loop(target int) int {
	if target <= 2 {
		return 1
	}
	return fibonacci_loop(target-1) + fibonacci_loop(target-2)
}
