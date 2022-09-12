package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var p = headA
	var q = headB
	for p != q {
		p = nextStep(p, headB)
		q = nextStep(q, headA)
	}
	return p
}

func nextStep(p, next *ListNode) *ListNode {
	if p == nil {
		return next
	}
	return p.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
