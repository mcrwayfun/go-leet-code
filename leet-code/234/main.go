package main

import "fmt"

/**
题目：
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

示例 1：
输入：head = [1,2,2,1]
输出：true

示例 2：
输入：head = [1,2]
输出：false

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func isPalindrome(head *ListNode) bool
*/

/**
解题思路1：
1：遍历链表，使用一个栈来存储访问过的元素
2：出栈的顺序就是逆序的，再次遍历链表，和栈内元素相比，若全部相同则返回true

time complexity: O(n)
space complexity: O(n)
*/
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	var stack []int
	var cur = head
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}

	cur = head
	for cur != nil && len(stack) > 0 {
		if stack[len(stack)-1] != cur.Val {
			return false
		}
		cur = cur.Next
		stack = stack[:len(stack)-1]
	}

	return true
}

/**
解法2：可以逆转链表来判断是否为回文链表
1：逆转一半的链表，然后将使用pre和cur来判断是否为回文链表
2：如果链表长度为奇数，则cur向前移动一步

time complexity: O(n)
space complexity: O(1)
*/
func isPalindromeV2(head *ListNode) bool {
	if head == nil {
		return false
	}

	var length int
	for p := head; p != nil; p = p.Next {
		length++
	}

	var cur = head
	var pre *ListNode
	for i := 0; i < length/2; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	if length%2 != 0 {
		cur = cur.Next
	}

	for cur != nil && pre != nil {
		if cur.Val != pre.Val {
			return false
		}
		cur = cur.Next
		pre = pre.Next
	}

	return true
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
