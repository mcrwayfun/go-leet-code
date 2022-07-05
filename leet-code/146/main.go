package main

/**
请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value)如果关键字key 已经存在，则变更其数据值value ；
如果不存在，则向缓存中插入该组key-value 。如果插入操作导致关键字数量超过capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。



示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/lru-cache
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type LRUCache struct {

}


func Constructor(capacity int) LRUCache {

}


func (this *LRUCache) Get(key int) int {

}


func (this *LRUCache) Put(key int, value int)  {

}
*/

type LRUCache struct {
	Head *Node
	Data map[int]*Node
}

type Node struct {
	Val  int
	Key  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{-1, -1, nil, nil}
	cur := head

	// 构造一个双向链表
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
		if c.Head.Val != -1 {
			delete(c.Data, c.Head.Key)
		}
		c.Head.Key = key
		c.Head.Val = value
		c.Data[key] = c.Head
		c.Head = c.Head.Prev
	}
}

func (c *LRUCache) move2head(cur *Node) {
	if cur == c.Head {
		c.Head = c.Head.Next
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
/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
