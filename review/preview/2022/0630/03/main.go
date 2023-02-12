package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var carry int
	var dummy = &ListNode{}
	var p = dummy

	for l1 != nil || l2 != nil || carry != 0 {
		var sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
		carry = sum % 10
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
