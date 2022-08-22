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
		p = nextStep(p, headB)
		q = nextStep(q, headA)
	}
	return p
}

func nextStep(p, q *ListNode) *ListNode {
	if p != nil {
		return p.Next
	}
	return q
}


type ListNode struct {
	Val  int
	Next *ListNode
}
