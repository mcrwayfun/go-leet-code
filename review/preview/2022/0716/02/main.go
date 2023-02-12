package main

type LRUCache struct {
	Head *Node
	Data map[int]*Node
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	cur := &Node{-1, -1, nil, nil}
	var head = cur
	for i := 0; i < capacity-1; i++ {
		cur.Next = &Node{-1, -1, cur, nil}
		cur = cur.Next
	}

	head.Prev = cur
	cur.Next = head
	return LRUCache{
		Head: head,
		Data: make(map[int]*Node),
	}
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
	if _, ok := c.Data[key]; ok {// 存在则是更新
		cur := c.Data[key]
		cur.Val = value
		c.move2head(cur)
		return
	} else {
		if c.Head.Val != -1 {
			delete(c.Data, c.Head.Key)
		}
		c.Head.Key = key
		c.Head.Val = value
		c.Data[key] = c.Head
		c.move2head(c.Head)
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

	cur.Prev = c.Head
	c.Head.Next.Prev = cur

	cur.Next = c.Head.Next
	c.Head.Next = cur
}
