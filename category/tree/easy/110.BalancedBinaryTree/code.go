package leetcode

import "math"

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

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHight := maxDepth(root.Left)
	rightHight := maxDepth(root.Right)
	return math.Abs(float64(leftHight-rightHight)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
