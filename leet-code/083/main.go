package main

/**
给定一个已排序的链表的头head，删除所有重复的元素，使每个元素只出现一次。返回 已排序的链表。

示例 1：
输入：head = [1,1,2]
输出：[1,2]

示例 2：
输入：head = [1,1,2,3,3]
输出：[1,2,3]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func deleteDuplicates(head *ListNode) *ListNode{}

type ListNode struct {
	Val  int
	Next *ListNode
}
 */

// time complexity: O(n)
// space complexity: O(n)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var cur = head
	var next = head
	for next != nil {
		if cur.Val == next.Val {// skip the same ele
			cur.Next = next.Next
		} else {
			cur = cur.Next
		}
		next = next.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

