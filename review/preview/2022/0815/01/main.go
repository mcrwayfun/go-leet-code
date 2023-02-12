package main

// time complexity: O(m+n)
// space complexity: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var p = headA
	var q = headB
	for p != q {
		p = next(p, headB)
		q = next(q, headA)
	}
	return p
}

func next(head *ListNode, next *ListNode) *ListNode {
	if head == nil {
		return next
	}
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}