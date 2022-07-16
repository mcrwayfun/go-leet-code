package main

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

	if n != 0 { // n长度大于链表
		return head
	}

	for p.Next != nil {
		p = p.Next
		q = q.Next
	}

	// remove the node
	q.Next = q.Next.Next
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
