package main

import (
	"fmt"
)

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]

示例 2：
输入：l1 = [], l2 = []
输出：[]

示例 3：
输入：l1 = [], l2 = [0]
输出：[0]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode{}

type ListNode struct {
	Val  int
	Next *ListNode
}
*/

/**
使用一个指针p来遍历链表l1和l2
1：如果L1.Val < l2.Val, 则
	p.Next = l1 then l1 = l1.Next
2：否则 p.Next = l2 then l2 = l2.Next
3：p = p.Next
4：循环结束后，l1和l2可能存在未遍历结束的情况，哪个不为空则直接
将其拼在p后面即可

time complexity: O(m+n)
space complexity: O(1)
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var dummy = &ListNode{}
	p := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}

	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return dummy.Next
}

func main() {

	head3 := &ListNode{4, nil}
	head2 := &ListNode{2, head3}
	head1 := &ListNode{1, head2}

	head6 := &ListNode{4, nil}
	head5 := &ListNode{3, head6}
	head4 := &ListNode{1, head5}

	newHead := mergeTwoLists(head1, head4)
	println(newHead.String())
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
