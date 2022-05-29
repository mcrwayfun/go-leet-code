package main

import "math"

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{}
}

/**
入栈的时候，
1：如果min为空或者 val<=GetMin, 则min.push(val)
2：stack.push(val)
*/
func (this *MinStack) Push(val int) {
	if len(this.min) == 0 || this.GetMin() >= val {
		this.min = append(this.min, val)
	}
	this.stack = append(this.stack, val)
}

/**
出栈的时候，
1：如果Top==GetMin, 则min也要出栈
2：stack.pop
*/
func (this *MinStack) Pop() {
	if this.Top() == this.GetMin() {
		this.min = this.min[:len(this.min)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) > 0 {
		return this.min[len(this.min) - 1]
	}
	return math.MaxInt64
}
