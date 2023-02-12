package main

// time complexity: O(n)
// space complexity: O(1)
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	var length int
	for p := head; p != nil; p = p.Next {
		length++
	}

	var second = head
	var first *ListNode
	// 遍历的过程将前半部分反转
	for i := 0; i < length/2; i++ {
		next := second.Next
		second.Next = first
		first = second
		second = next
	}

	// 奇数需要向前走一步
	if length%2 == 1 {
		second = second.Next
	}

	for first != nil && second != nil {
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}

	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}
