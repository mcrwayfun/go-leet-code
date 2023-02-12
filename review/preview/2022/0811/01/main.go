package main

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

type LRUCache struct {
	Head *Node
	Data map[int]*Node
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
	return LRUCache{
		Head: head,
		Data: make(map[int]*Node),
	}
}

func (c *LRUCache) move2head(cur *Node) {
	if c.Head == cur {
		c.Head = c.Head.Prev
		return
	}

	// detach
	cur.Next.Prev = cur.Prev
	cur.Prev.Next = cur.Next

	// attach
	cur.Prev = c.Head
	cur.Next = c.Head.Next
	c.Head.Next.Prev = cur
	c.Head.Next = cur
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.Data[key]; !ok {
		return -1
	}
	cur := c.Data[key]
	c.move2head(cur)
	return cur.Val
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.Data[key]; ok {
		cur := c.Data[key]
		cur.Val = value
		c.move2head(cur)
	} else {
		if c.Head.Key != -1 {
			delete(c.Data, c.Head.Key)
		}
		c.Head.Key = key
		c.Head.Val = value
		c.Data[key] = c.Head
		c.move2head(c.Head)
	}
}
