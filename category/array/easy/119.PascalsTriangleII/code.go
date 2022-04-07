package leetcode

//可以使用DP的演算法

/*
	Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
	如果從0開始算起 第n 行的共有 n + 1 個數量,

*/

//第一種 全部找到後輸出
func getRow_First(rowIndex int) []int {
	row := make([][]int, rowIndex+1)
	for i := range row {
		row[i] = make([]int, i+1)
		row[i][0], row[i][i] = 1, 1
		for j := 1; j < i; j++ {
			row[i][j] = row[i-1][j-1] + row[i-1][j]
		}
	}
	return row[rowIndex]
}

//第二種 運用動態運算DP來輸出
func getRow_Sec(rowIndex int) []int {
	var pre, cur []int
	//層數
	for i := 0; i <= rowIndex; i++ {
		cur = make([]int, i+1) //第n行有i+1個數
		cur[0], cur[i] = 1, 1
		for j := 1; j < i; j++ {
			cur[j] = pre[j-1] + pre[j] //當前的第 j 個數 是上一個 j-1數 + 上一個 j 數
		}
		pre = cur
	}
	return pre
}

/*
	可以搞出推定 C i取n =  C i取n-1 + C i-1取n-1


*/
func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i; j > 0; j-- {
			row[j] += row[j-1]
		}
	}
	return row
}
