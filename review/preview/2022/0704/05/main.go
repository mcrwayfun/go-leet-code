package main

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	var prev = &ListNode{Next: head}
	var p = prev

	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}

	return prev.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
