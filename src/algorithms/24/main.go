package main

import "fmt"

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummyHead := &ListNode{Next: head}
	pre := dummyHead

	for pre.Next != nil && pre.Next.Next != nil {
		node1 := pre.Next
		node2 := node1.Next
		next := node2.Next

		pre.Next = node2
		node1.Next = next
		node2.Next = node1

		pre = node1
	}

	return dummyHead.Next
}

func main(){
	head4 := &ListNode{4, nil}
	head3 := &ListNode{3, head4}
	head2 := &ListNode{2, head3}
	head1 := &ListNode{1, head2}

	newHead := swapPairs(head1)
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