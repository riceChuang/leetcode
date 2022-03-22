package leetcode

import (
	"fmt"
	"strings"
)

/*
 主要用來製作 slice給出來的直變成一顆書
 一招此範例會跑出
     12
   /   \
  5     9
 / \   / \
6   2 0   8
   / \
  7  4
*/

type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

func test() {
	strArr := []string{"[12, 5, 9, 6, 2, 0, 8, #, #, 7, 4, #, #, #, #]", "6", "4"}
	question := strArr[0]
	question = strings.Replace(question, " ", "", -1)
	question = strings.Replace(question, "[", "", -1)
	question = strings.Replace(question, "]", "", -1)
	str := strings.Split(question, ",")
	fmt.Println(str)
	allNodes, root := makeTree(str)
	fmt.Println(allNodes)
	fmt.Println(root)
}

func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := lowestCommonAncestor236(root.Left, p, q)
	right := lowestCommonAncestor236(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}

func makeTree(strArr []string) ([]*TreeNode, *TreeNode) {
	var root *TreeNode
	var allNodes = []*TreeNode{}
	for i, str := range strArr {
		node := &TreeNode{
			Value: str,
		}
		if i == 0 {
			root = node
		}
		allNodes = append(allNodes, node)
	}

	leftAdd := 1
	rightAdd := 2
	for i, node := range allNodes {
		if (i+leftAdd) > len(allNodes) || (i+rightAdd) > len(allNodes) {
			break
		}

		if allNodes[i+leftAdd].Value != "#" {
			node.Left = allNodes[i+leftAdd]
		}

		if allNodes[i+rightAdd].Value != "#" {
			node.Right = allNodes[i+rightAdd]
		}

		leftAdd++
		rightAdd++
	}
	return allNodes, root
}
