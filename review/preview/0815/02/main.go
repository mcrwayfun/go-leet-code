package main

// time complexity: O(n)
// space complexity: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	var prev = &ListNode{Next: head}
	var dummy = prev
	for dummy.Next != nil {
		if dummy.Next.Val == val {
			dummy.Next = dummy.Next.Next
		} else {
			dummy = dummy.Next
		}
	}
	return prev.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
