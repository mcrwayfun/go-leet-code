package main

// time complexity: O(n)
// space complexity: O(1)
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	var length int
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	var cur = head
	var prev *ListNode

	for i := 0; i < length/2; i++ {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	if length&1 == 1 { // 奇数向前挪动一个位置
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

// time complexity: O(n)
// space complexity: O(n)
func isPalindromeStack(head *ListNode) bool {
	if head == nil {
		return false
	}

	var stack = make([]*ListNode, 0)
	for p := head; p != nil; p = p.Next {
		stack = append(stack, p)
	}

	for p := head; p != nil; p = p.Next {
		if p.Val != stack[len(stack)-1].Val {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
