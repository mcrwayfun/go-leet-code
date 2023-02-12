package main

/**
定义dummy节点，使用两个指针p和q。
1：指针p先走n步，同时n--
2：如果n!=0，则不需要移除链表节点，因为n比链表长度大
3：如果p.Next!=nil, 则p和q不断向前移动
4：q.Next = q.Next.Next (删除倒数第N个节点）
*/
// time complexity: O(n)
// space complexity: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head
	var p = dummy
	var q = dummy
	for n > 0 && p.Next != nil {
		p = p.Next
		n--
	}
	if n != 0 {
		return dummy.Next
	}

	for p.Next != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
