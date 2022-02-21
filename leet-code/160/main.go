package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var aLen, bLen int
	curA, curB := headA, headB
	for curA != nil {
		aLen++
		curA = curA.Next
	}

	for curB != nil {
		bLen++
		curB = curB.Next
	}

	var sub int
	curA, curB = headA, headB
	if aLen > bLen {
		sub = aLen - bLen
		for curA != nil && sub != 0 {
			curA = curA.Next
			sub--
		}
	} else {
		sub = bLen - aLen
		for curB != nil && sub != 0 {
			curB = curB.Next
			sub--
		}
	}

	if curA == curB {
		return curA
	}

	for curA != nil && curA.Next != nil && curB != nil && curB.Next != nil {
		if curA.Next == curB.Next {
			return curA.Next
		}
		curA = curA.Next
		curB = curB.Next
	}

	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	curA, curB := headA, headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}

		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}

	return curA
}
