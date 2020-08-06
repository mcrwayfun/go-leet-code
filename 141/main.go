package main

import "fmt"

func hasCycle(head *ListNode) bool {
	first,second := head,head
	for first != nil && first.Next != nil{
		first = first.Next.Next
		second = second.Next
		if first == second {
			return true
		}
	}
	return false
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