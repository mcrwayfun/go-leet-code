package main

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	var length int
	for p := head; p != nil; p = p.Next {
		length++
	}

	var prev = &ListNode{}
	var cur = head
	for i := 0; i < length/2; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	if length%2 != 0 {
		cur = cur.Next
	}

	for cur != nil && prev != nil {
		if cur.Val != prev.Val {
			return false
		}
		cur = cur.Next
		prev = prev.Next
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
