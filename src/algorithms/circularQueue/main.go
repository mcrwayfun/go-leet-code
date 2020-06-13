package main

import (
	"algorithms/circularQueue/queue"
	"fmt"
)

func main() {
	myCircularQueue := queue.Constructor(3)
	fmt.Println(myCircularQueue.EnQueue(1))
	fmt.Println(myCircularQueue.EnQueue(2))
	fmt.Println(myCircularQueue.EnQueue(3))
	fmt.Println(myCircularQueue.EnQueue(4))
	fmt.Println(myCircularQueue.Rear())
	fmt.Println(myCircularQueue.IsFull())
	fmt.Println(myCircularQueue.DeQueue())
	fmt.Println(myCircularQueue.EnQueue(4))
	fmt.Println(myCircularQueue.Rear())
}
