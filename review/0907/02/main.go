package main

// time complexity: O(n)
// space complexity: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	var dummy = &ListNode{Next: head}
	var prev = dummy
	for prev.Next != nil {
		if prev.Next.Val == val {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
