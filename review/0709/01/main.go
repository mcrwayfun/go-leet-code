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
	var head = &Node{-1, -1, nil, nil}
	var cur = head

	// 构造双向的循环链表
	for i := 0; i < capacity; i++ {
		cur.Next = &Node{-1, -1, cur, nil}
		cur = cur.Next
	}

	// 头尾指针相连
	cur.Next = head
	head.Prev = cur
	return LRUCache{Head: head, Data: make(map[int]*Node)}
}

func (c *LRUCache) Get(key int) int {
	if _, ok := c.Data[key]; !ok {
		return -1
	}
	cur := c.Data[key]
	c.moveToHead(cur)
	return cur.Val
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.Data[key]; ok {
		// 存在则进行覆盖
		cur := c.Data[key]
		cur.Val = value
		c.moveToHead(cur)
	} else { // 不存在则新增
		if c.Head.Val != -1 {
			// head当前存在元素，需要移除
			delete(c.Data, c.Head.Key)
		}

		c.Head.Key = key
		c.Head.Val = value
		c.Data[key] = c.Head
		c.Head = c.Head.Prev
	}
}

func (c *LRUCache) moveToHead(cur *Node) {
	if c.Head == cur {
		c.Head = c.Head.Prev // 向前移动到当前最活跃的位置
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
