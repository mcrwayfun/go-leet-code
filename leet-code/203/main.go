package main

/**
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。

示例 1：
输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]

示例 2：
输入：head = [], val = 1
输出：[]

示例 3：
输入：head = [7,7,7,7], val = 7
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-linked-list-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func removeElements(head *ListNode, val int) *ListNode {}

type ListNode struct {
	Val  int
	Next *ListNode
}
*/

// time complexity: O(n)
// space complexity: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	notEqual := dummy
	for notEqual.Next != nil {
		if notEqual.Next.Val == val {
			notEqual.Next = notEqual.Next.Next
		} else {
			notEqual = notEqual.Next
		}
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
