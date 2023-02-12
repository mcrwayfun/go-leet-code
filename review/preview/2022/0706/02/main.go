package main

import "fmt"

type LRUCache struct {
	Head *Node
	Data map[int]*Node
}

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

func Constructor(capacity int) LRUCache {
	var head = &Node{-1, -1, nil, nil}
	var cur = head
	for i := 0; i < capacity-1; i++ {
		cur.Next = &Node{-1, -1, cur, nil}
		cur = cur.Next
	}

	cur.Next = head
	head.Prev = cur
	return LRUCache{head, make(map[int]*Node)}
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.Data[key]; !ok {
		return -1
	}
	cur := c.Data[key]
	c.move2head(cur)
	return cur.Value
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.Data[key]; ok {
		cur := c.Data[key]
		cur.Value = value
		c.move2head(cur)
	} else {
		if c.Head.Key != -1 {
			delete(c.Data, c.Head.Key)
		}
		c.Head.Key = key
		c.Head.Value = value
		c.Data[key] = c.Head
		c.Head = c.Head.Prev
	}
}

func (c *LRUCache) move2head(cur *Node) {
	if c.Head == cur {
		c.Head = c.Head.Prev
		return
	}

	// detach
	cur.Prev.Next = cur.Next
	cur.Next.Prev = cur.Prev

	// attach
	cur.Next = c.Head.Next
	cur.Next.Prev = cur
	c.Head.Next = cur
	cur.Prev = c.Head
}

func main() {
	/**
	["LRUCache","put","put","get","put","get","put","get","get","get"]
	[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	*/
	//obj := Constructor(2)
	//obj.Put(1, 1)
	//obj.Put(2, 2)
	//fmt.Println(obj.Get(1))
	//obj.Put(3, 3)
	//fmt.Println(obj.Get(2))
	//obj.Put(4, 4)
	//fmt.Println(obj.Get(1))
	//fmt.Println(obj.Get(3))
	//fmt.Println(obj.Get(4))

	/**
	["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
	[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]
	*/
	obj := Constructor(3)
	obj.Put(1, 1)
	obj.Put(2, 2)
	obj.Put(3, 3)
	obj.Put(4, 4)
	fmt.Println(obj.Get(4))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(1))
	obj.Put(5, 5)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
	fmt.Println(obj.Get(5))
}
