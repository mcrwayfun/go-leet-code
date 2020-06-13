package queue

type MyCircularQueue struct {
	capacity int
	data     []int
	front    int // 队头下标
	rear     int // 队尾下标
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		capacity: k+1,
		data:     make([]int, k+1),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.rear] = value
	// 判断rear是否需要调头
	this.rear = (this.rear + 1) % this.capacity
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	// 判断front是否需要调头
	this.front = (this.front + 1) % this.capacity
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	// 该方法可被外部调用，需要判断队列是否为空
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	// 该方法可被外部调用，需要判断队列是否为空
	if this.IsEmpty() {
		return -1
	}
	// 获取当前的
	return this.data[(this.rear-1+this.capacity)%this.capacity]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.front == (this.rear+1)%this.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
