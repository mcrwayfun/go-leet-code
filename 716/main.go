package main

import (
	"fmt"
	"math"
)

type MaxStack struct {
	stack []int
	maxStack []int
}


/** initialize your data structure here. */
func Constructor() MaxStack {
	return MaxStack{
		stack: []int{},
		maxStack: []int{math.MinInt64},
	}
}


func (this *MaxStack) Push(x int)  {
	this.stack = append(this.stack, x)
	if this.PeekMax() < x {
		this.maxStack = append(this.maxStack, x)
	} else {
		this.maxStack = append(this.maxStack, this.PeekMax())
	}
}


func (this *MaxStack) Pop() int {
	ret := this.stack[len(this.stack) - 1]
	this.stack = this.stack[:len(this.stack) - 1]
	this.maxStack = this.maxStack[:len(this.maxStack) - 1]
	return ret
}


func (this *MaxStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MaxStack) PeekMax() int {
	return this.maxStack[len(this.maxStack) - 1]
}

func (this *MaxStack) PopMax() int {
	var tmpStack []int // 临时栈
	max := this.PeekMax()
	for this.Top() != max {
		tmpStack = append(tmpStack, this.Pop())
	}
	this.Pop() // 删除当前栈的最大元素
	for i:=len(tmpStack)-1;i>=0;i--{
		this.Push(tmpStack[i]) // 将元素重新压回stack
	}

	return max
}

func main(){
	testCase2()
}

func testCase1(){
	/**
	["MaxStack","push","push","push","top","popMax","top","peekMax","pop","top"]
	[[],[5],[1],[5],[],[],[],[],[],[]]
	*/
	maxStack := Constructor()
	maxStack.Push(5)
	maxStack.Push(1)
	maxStack.Push(5)
	fmt.Printf("maxStack after push 5->1->5: %v\n", maxStack)

	top := maxStack.Top()
	fmt.Printf("maxStack after top %d, maxStack: %v\n", top, maxStack)

	max := maxStack.PopMax()
	fmt.Printf("maxStack after popMax %d, maxStack: %v\n", max, maxStack)

	top = maxStack.Top()
	fmt.Printf("maxStack after top %d, maxStack: %v\n", top, maxStack)

	peekMax := maxStack.PeekMax()
	fmt.Printf("maxStack after peekMax %d, maxStack: %v\n", peekMax, maxStack)

	pop := maxStack.Pop()
	fmt.Printf("maxStack after pop %d, maxStack: %v\n", pop, maxStack)

	top = maxStack.Top()
	fmt.Printf("maxStack after top %d, maxStack: %v\n", top, maxStack)
}

func testCase2(){
	/**
	["MaxStack","push","push","popMax","peekMax"]
	[[],[5],[1],[],[]]
	 */
	maxStack := Constructor()
	maxStack.Push(5)
	maxStack.Push(1)
	fmt.Printf("maxStack after push 5->1->5: %v\n", maxStack)

	max := maxStack.PopMax()
	fmt.Printf("maxStack after popMax %d, maxStack: %v\n", max, maxStack)

	peekMax := maxStack.PeekMax()
	fmt.Printf("maxStack after peekMax %d, maxStack: %v\n", peekMax, maxStack)
}

