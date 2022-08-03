package main

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
	Val int
	Next *ListNode
}
