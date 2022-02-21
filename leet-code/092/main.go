package main

import "fmt"

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	changeLen := n - m + 1 // 逆置的长度
	var preHead *ListNode  // 逆置头节点的前驱
	cur := head
	result := head

	// 1:遍历到要逆置的节点
	for cur != nil && m > 1 {
		preHead = cur
		cur = cur.Next
		m--
	}

	newTail := cur        // 需要逆置的头节点，逆置后的尾节点
	var newHead *ListNode // 需要逆置的尾节点，逆置后的头节点

	// 2:逆置m-n的节点
	for cur != nil && changeLen != 0 {
		tmp := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = tmp
		changeLen--
	}

	// 3:尾节点
	if newTail != nil {
		newTail.Next = cur
	}

	if preHead != nil {
		preHead.Next = newHead
	} else {
		result = newHead
	}

	return result
}

func reverseBetween3(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}

	a, b := m, n-m+1
	var prev *ListNode        // 保留逆转前的节点prev
	cur, result := head, head // 返回的结果

	// 遍历到逆置节点
	for i := 0; i < a-1; i++ {
		prev = cur
		cur = cur.Next
	}

	var newHead *ListNode // 逆置新的头节点
	newTail := cur
	for i := 0; i < b; i++ {
		tmp := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = tmp
	}

	if prev != nil { // m != 1
		prev.Next = newHead
	} else {
		result = newHead
	}
	newTail.Next = cur
	return result
}

func main() {
	head5 := &ListNode{5, nil}
	head4 := &ListNode{4, head5}
	head3 := &ListNode{3, head4}
	head2 := &ListNode{2, head3}
	head1 := &ListNode{1, head2}

	newHead := reverseBetween3(head1, 2, 4)
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
