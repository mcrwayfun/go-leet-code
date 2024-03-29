package main

/**
给你一个链表的头节点 head ，判断链表中是否有环。
如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
注意：pos 不作为参数进行传递。仅仅是为了标识链表的实际情况。如果链表中存在环，则返回 true 。 否则，返回 false 。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。

示例2：
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：false
解释：链表中没有环。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/linked-list-cycle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func hasCycle(head *ListNode) bool {}

type ListNode struct {
	Val  int
	Next *ListNode
}
 */

/**
使用快慢指针，快指针每次走2步，慢指针每次走一步，如果有环，
则快慢指针必定在环上相遇（可以画个例子验证下）

现在来验证这道题的时间复杂度为何为O(n)
1：假设链表有n个节点，环上有k个节点，则非环节点有n-k个
2：慢指针在入环前需要走n-k步，假设快指针和慢指针之间的距离节点为d，
显然有d<k(因为此时快指针还在环上)
3：快慢指针之间的速度差为1，慢指针需要走d次才能和快指针相遇，因此慢指针
在环内走了d次。所以走的节点为：n-k+d
4：由于d<k，所以有n-k+d<n

time complexity: O(n)
space complexity: O(1)
 */
func hasCycle(head *ListNode) bool {
	var fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

type ListNode struct {
	Val  int
	Next *ListNode
}

