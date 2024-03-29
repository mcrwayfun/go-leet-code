package main

/**
题目：
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，
并且每个节点只能存储一位数字。请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0开头。

示例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]

示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{}

type ListNode struct {
	Val int
	Next *ListNode
}
*/

// time complexity: O(max(m,n))
// space complexity: O(max(m,n))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	var carry int

	for l1 != nil || l2 != nil || carry != 0 {
		var sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
		carry = sum / 10
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
