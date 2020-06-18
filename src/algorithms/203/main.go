package main

import "fmt"

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummyNode := &ListNode{Next: head}
	cur := dummyNode
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}else{
			cur = cur.Next
		}
	}

	return dummyNode.Next
}

func main(){
	head5 := &ListNode{5, nil}
	head4 := &ListNode{4, head5}
	head3 := &ListNode{3, head4}
	head2 := &ListNode{2, head3}
	head1 := &ListNode{1, head2}

	elements := removeElements(head1, 3)
	println(elements.String())
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