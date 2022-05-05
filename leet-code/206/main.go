package main

/**
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：
输入：head = [1,2]
输出：[2,1]

示例 3：
输入：head = []
输出：[]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func reverseList(head *ListNode) *ListNode
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
解法：反转链表并返回反转后的头节点：需要维护三个变量，一个是指向当前节点的cur，
一个是指向当前节点的先前节点的pre，一个是指向当前节点的下一节点的next。

举个例子，比如有 1->2->3->4->5->null
pre=null, cur=1, next=2
1：先记录当前节点的next关系，因为反转后这个关系就不存在了。next = cur.Next
2：当前节点指向的下一节点反转，cur.Next = pre
3：pre移动到cur的位置，pre = cur
4：cur移动到next的位置，cur = next
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var cur = head
	var pre *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
