package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	var cur = head
	var k int
	for cur != nil {
		cur = cur.Next
		k++
	}

	if k-n < 0 {
		return nil
	}

	var dummy = &ListNode{}
	var prev = dummy
	dummy.Next = head
	for i := 0; i < k-n; i++ {
		dummy = dummy.Next
	}

	dummy.Next = dummy.Next.Next

	return prev.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
