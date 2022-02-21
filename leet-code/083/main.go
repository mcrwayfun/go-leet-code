package main

import "fmt"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		}
		cur = cur.Next
	}
	return head
}

func main() {
	head5 := &ListNode{3, nil}
	head4 := &ListNode{3, head5}
	head3 := &ListNode{2, head4}
	head2 := &ListNode{1, head3}
	head1 := &ListNode{1, head2}

	newHead := deleteDuplicates(head1)
	println(newHead.String())
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
