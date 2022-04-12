package leetcode

import (
	"fmt"
)

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
	}
}

func StackFunc() {
	sk := Constructor()
	sk.Push(2)
	sk.Push(0)
	sk.Push(3)
	sk.Push(0)
	fmt.Println(sk.GetMin())
	sk.Pop()
	fmt.Println(sk.GetMin())
	sk.Pop()
	fmt.Println(sk.GetMin())
	sk.Pop()
	fmt.Println(sk.GetMin())

}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) > 0 {
		return this.stack[len(this.stack)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	var result int
	for i, value := range this.stack {
		if i == 0 {
			result = value
			continue
		}
		if result > value {
			result = value
		}
	}
	return result
}
