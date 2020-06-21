package main

import "fmt"

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	values := make([]int, 0)
	cur := head
	for cur != nil {
		values = append(values, cur.Val)
		cur = cur.Next
	}

	count := len(values)
	for i := 0; i < count/2; i++ {
		if values[i] != values[count-1-i] {
			return false
		}
	}

	return true
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	cur := l
	var str string
	for cur != nil {
		str = fmt.Sprintf("%s %d", str, cur.Val)
		cur = cur.Next
	}
	return str
}
