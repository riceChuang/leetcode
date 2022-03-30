package leetcode

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

//遍歷所有節點
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
		tempLeft := preorderTraversal(root.Left)
		res = append(res, tempLeft...)
		tempRight := preorderTraversal(root.Right)
		res = append(res, tempRight...)
	}
	return res
}
