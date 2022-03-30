package leetcode

import "fmt"

/*
 num & 1 的概念為
 假如 num binary = 11111 , 1 的binary = 00001
 11111 & 00001 = 1
 假如 num binary = 10100 , 1 的binary = 00001
 10100 & 00001 = 0
 可以拿到最後一位為1 or 0
*/

func reverseBits(num uint32) uint32 {
	var res uint32
	fmt.Printf("question:%b", num)
	for i := 0; i < 32; i++ {
		fmt.Printf("res:%b, data:%v \n", res, num&1)
		res = res<<1 | num&1
		num >>= 1
	}
	return res
}
