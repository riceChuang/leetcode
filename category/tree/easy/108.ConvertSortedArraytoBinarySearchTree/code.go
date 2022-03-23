package leetcode

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	//nums 目前是升序的
	sort.Ints(nums)
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBST(nums[:len(nums)/2]),
		Right: sortedArrayToBST(nums[len(nums)/2+1:])}
}

/*
	測試數據 [ -10, -3 , 0 , 5 , 9 ]

						0
		左邊[-10 , -3, 0]      右邊 [5,9]
			-3					9
	左邊[-10]   右邊[]		左邊[5]  右邊[]
		10
*/