package main

import "fmt"

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

// 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	// 新的头结点
	var newHead *ListNode = nil
	cur := head
	for cur != nil {
		// 1:备份下一个结点
		tmp := cur.Next
		// 2:当前结点指向新的头结点
		cur.Next = newHead
		// 3:新的头结点移动到当前结点
		newHead = cur
		// 4:head结点移动
		cur = tmp
	}

	return newHead
}

// 反转链表
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 反转后的头指针
	ret := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

func main() {
	tail := &ListNode{-1, nil}
	head5 := &ListNode{5, tail}
	head4 := &ListNode{4, head5}
	head3 := &ListNode{3, head4}
	head2 := &ListNode{2, head3}
	head1 := &ListNode{1, head2}

	newHead := reverseList(head1)
	println(newHead.String())
}
