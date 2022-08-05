package main

/**
给定一个已排序的链表的头head ，删除原始链表中所有重复数字的节点，只留下不同的数字。返回 已排序的链表。

示例 1：
输入：head = [1,2,3,3,4,4,5]
输出：[1,2,5]

示例 2：
输入：head = [1,1,1,2,3]
输出：[2,3]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func deleteDuplicates(head *ListNode) *ListNode {}

type ListNode struct {
	Val  int
	Next *ListNode
}
 */

// time complexity: O(n)
// space complexity: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var dummy = &ListNode{Next: head}
	var prev = dummy
	var cur = dummy.Next
	for cur != nil {
		// cur 移动到最后一个重复元素
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		// prev 和 cur没有相隔，则需要更新prev.Next到重复元素的下一个
		if prev.Next != cur {
			prev.Next = cur.Next
		} else {
			prev = prev.Next // 两者相邻，否则prev移动
		}
		cur = prev.Next
	}
	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}


