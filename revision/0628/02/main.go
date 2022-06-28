package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p := headA
	q := headB

	for p != q {
		p = next(p, headB)
		q = next(q, headA)
	}
	return p
}

func next(head *ListNode, new *ListNode) *ListNode {
	if head == nil {
		return new
	}
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
