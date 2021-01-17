package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	carry := 0
	cur := &ListNode{}
	ret := cur
	for l1 != nil || l2 != nil {
		x := getVal(l1)
		y := getVal(l2)

		cur.Next = &ListNode{Val: (x+y+carry)%10}
		cur = cur.Next
		carry = (x+y+carry)/10 // 获取进位

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {// 最后一位仍有进位
		cur.Next = &ListNode{Val: carry}
	}

	return ret.Next
}

func getVal(head *ListNode) int {
	if head != nil {
		return head.Val
	} else {
		return 0
	}
}

func main(){
	l1 := NewListNodes([]int{2,4,3})
	l2 := NewListNodes([]int{5,6,4})

	numbers := addTwoNumbers(l1, l2)
	fmt.Println(numbers)
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
