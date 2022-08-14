package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var p, q = headA, headB
	for p != q {
		p = next(p, headB)
		q = next(q, headA)
	}
	return p
}

func next(headA, headB *ListNode) *ListNode {
	if headA != nil {
		return headA.Next
	}
	return headB
}

type ListNode struct {
	Val  int
	Next *ListNode
}
