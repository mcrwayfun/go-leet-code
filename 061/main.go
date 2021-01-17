package main

import "fmt"

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	l := getLen(head)
	k = k % l

	if k == 0 {
		return head
	}

	// 遍历到偏移点
	end := head
	for i := 0; i < k; i++ {
		end = end.Next
	}

	start := head
	for end.Next != nil {
		start = start.Next
		end = end.Next
	}

	end.Next = head
	newHead := start.Next
	start.Next = nil
	return newHead
}

func getLen(head *ListNode) int {
	res := 0
	cur := head
	for cur != nil {
		res++
		cur = cur.Next
	}
	return res
}

func main() {
	head := NewListNodes([]int{1, 2, 3, 4, 5})
	right := rotateRight(head, 4)
	fmt.Printf("%s", right)
}

func NewListNodes(data []int) *ListNode {
	var head = &ListNode{}
	var prev = head
	for _, v := range data {
		head.Next = &ListNode{Val: v}
		head = head.Next
	}
	return prev.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var res string
	for l != nil {
		res = res + fmt.Sprintf("%d->", l.Val)
		l = l.Next
	}
	res = res + "NULL"
	return res
}
