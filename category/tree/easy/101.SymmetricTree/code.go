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

func isSymmetric(root *TreeNode) bool {
	return isSameTree(InvertTree(root.Left), root.Right)
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}

}

func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	InvertTree(root.Left)
	InvertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
