package main

// time complexity: O(n)
// space complexity: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var p, q = headA, headB
	for p != q {
		p = next(p, headB)
		q = next(q, headA)
	}
	return p
}

func next(p, q *ListNode) *ListNode {
	if p != nil {
		return p.Next
	}
	return q
}

type ListNode struct {
	Val  int
	Next *ListNode
}
