package main

type MinStack struct {
	stack []int
	min []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int)  {
	if len(this.min) == 0 || this.GetMin() >= val {
		this.min = append(this.min, val)
	}
	this.stack = append(this.stack, val)
}


func (this *MinStack) Pop()  {
	if this.GetMin() == this.Top() {
		this.min = this.min[:len(this.min) - 1]
	}
	this.stack = this.stack[:len(this.stack) - 1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack) - 1]
}


func (this *MinStack) GetMin() int {
	return this.min[len(this.min) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
