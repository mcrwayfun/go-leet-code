package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	dumpHead := &ListNode{} // 构建虚拟头节点
	cur := dumpHead
	carry := 0 // 进位

	for l1 != nil || l2 != nil {
		var a, b int
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}

		cur.Next = &ListNode{Val: (a + b + carry) % 10}
		cur = cur.Next

		carry = (a + b + carry) / 10 // 获取进位，题目说进位可能大于1
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry > 0 {// 最后一位为进位
		cur.Next = &ListNode{Val: carry % 10}
	}

	return dumpHead.Next
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
