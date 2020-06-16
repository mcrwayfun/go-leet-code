package main

import "fmt"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val { // 存在值相等的节点
			removeVal := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == removeVal {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}

func main(){
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
