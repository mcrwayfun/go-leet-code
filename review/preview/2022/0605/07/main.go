package main

// time complexity: O(n)
// space complexity: O(1)
func hasCycle(head *ListNode) bool {
	var slow, fast = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}
