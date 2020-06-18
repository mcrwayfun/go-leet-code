package main

import "fmt"

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head} // 虚拟头节点
	cur := dummy
	temp1 := &ListNode{}
	temp2 := temp1

	for cur.Next != nil {
		if cur.Next.Val > x {
			t := cur.Next
			cur.Next = cur.Next.Next
			temp2.Next = t
			temp2 = temp2.Next
		} else {
			cur = cur.Next
		}
	}

	temp2.Next = nil
	cur.Next = temp1.Next
	return dummy.Next
}

func partition2(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head} // 虚拟头节点
	dummy2 := dummy

	tempHead := &ListNode{}
	tempHead2 := tempHead

	for dummy2.Next != nil {
		if dummy2.Next.Val < x {
			dummy2 = dummy2.Next
		} else {
			t := dummy2.Next
			dummy2.Next = dummy2.Next.Next
			tempHead2.Next = t
			tempHead2 = tempHead2.Next
		}
	}

	tempHead2.Next = nil
	dummy2.Next = tempHead.Next
	return dummy.Next
}

func main() {
	head6 := &ListNode{2, nil}
	head5 := &ListNode{5, head6}
	head4 := &ListNode{2, head5}
	head3 := &ListNode{3, head4}
	head2 := &ListNode{4, head3}
	head1 := &ListNode{1, head2}

	newHead := partition2(head1,3)
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
