package main

// time complexity: O(n)
// space complexity: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	var prev = &ListNode{Next: head}
	var notEqual = prev
	for notEqual.Next != nil {
		if notEqual.Next.Val == val {
			notEqual.Next = notEqual.Next.Next
		} else {
			notEqual = notEqual.Next
		}
	}
	return prev.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
