package main

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 获取链表的中间节点
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 前后断开
	prev.Next = nil
	left := sortList(head)
	right := sortList(slow)

	return merge(left, right)
}

// 两个链表排序
func merge(left *ListNode, right *ListNode) *ListNode {
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	if left.Val < right.Val {
		left.Next = merge(left.Next, right)
		return left
	}
	right.Next = merge(left, right.Next)
	return right
}

type ListNode struct {
	Val  int
	Next *ListNode
}
