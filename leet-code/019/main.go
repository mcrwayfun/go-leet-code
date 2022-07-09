package main

/**
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

示例 1：
输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]

示例 2：
输入：head = [1], n = 1
输出：[]

示例 3：
输入：head = [1,2], n = 1
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：
func removeNthFromEnd(head *ListNode, n int) *ListNode {}

type ListNode struct {
	Val int
	Next *ListNode
}
*/

/**
定义dummy节点，使用两个指针p和q。
1：指针p先走n步，同时n--
2：如果n!=0，则不需要移除链表节点，因为n比链表长度大
3：如果p.Next!=nil, 则p和q不断向前移动
4：q.Next = q.Next.Next (删除倒数第N个节点）
*/
// time complexity: O(n)
// space complexity: O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummy = &ListNode{}
	dummy.Next = head
	var p = dummy
	var q = dummy
	for n > 0 && p.Next != nil {
		p = p.Next
		n--
	}
	if n != 0 {
		return dummy.Next
	}

	for p.Next != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
