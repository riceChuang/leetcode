package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法一 单链表
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	res := true
	// 寻找中间结点
	p1 := head
	p2 := head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}
	// 反转链表后半部分  1->2->3->4->5->6 to 1->2->3->6->5->4
	preMiddle := p1
	preCurrent := p1.Next
	for preCurrent.Next != nil {
		current := preCurrent.Next
		preCurrent.Next = current.Next
		current.Next = preMiddle.Next
		preMiddle.Next = current
	}
	// 扫描表，判断是否是回文
	p1 = head
	p2 = preMiddle.Next
	// fmt.Printf("p1 = %v p2 = %v preMiddle = %v head = %v\n", p1.Val, p2.Val, preMiddle.Val, L2ss(head))
	for p1 != preMiddle {
		// fmt.Printf("*****p1 = %v p2 = %v preMiddle = %v head = %v\n", p1, p2, preMiddle, L2ss(head))
		if p1.Val == p2.Val {
			p1 = p1.Next
			p2 = p2.Next
			// fmt.Printf("-------p1 = %v p2 = %v preMiddle = %v head = %v\n", p1, p2, preMiddle, L2ss(head))
		} else {
			res = false
			break
		}
	}
	if p1 == preMiddle {
		if p2 != nil && p1.Val != p2.Val {
			return false
		}
	}
	return res
}
