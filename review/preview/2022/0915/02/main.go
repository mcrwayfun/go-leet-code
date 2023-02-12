package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var p = headA
	var q = headB
	for p != q {
		p = nextListNode(p, headB)
		q = nextListNode(q, headA)
	}
	return p
}

func nextListNode(cur, next *ListNode) *ListNode {
	if cur != nil {
		return cur.Next
	}
	return next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
