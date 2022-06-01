package main

/**
给你两个单链表的头节点headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表不存在相交节点，返回 null 。

图示两个链表在节点 c1 开始相交：

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/intersection-of-two-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

模板：func getIntersectionNode(headA, headB *ListNode) *ListNode {}

type ListNode struct {
	Val  int
	Next *ListNode
}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
1：求两个链表的长度，
2：if lenA>lenB, headA向前走lenA-lenB
	else headB向前走lenB-lenA
3：headA和headB一起向前走，最后返回headA

time complexity: O(m+n)
space complexity: O(1)
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	var lenA, lenB int
	for p := headA; p != nil; p = p.Next {
		lenA++
	}
	for p := headB; p != nil; p = p.Next {
		lenB++
	}

	var p, q = headA, headB
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			p = p.Next
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			q = q.Next
		}
	}

	for p != q {
		p = p.Next
		q = q.Next
	}
	return p
}

/**
这道一道求交叉点的问题，如果从两个链表的开头开始，最终如果在某个节点相遇，那么
会走过相同的路程，由此可以有：
1：假设p从headA开始，如果走到末尾则到headB
2：q从headB开始，如果走到末尾则到headA
直到两者相遇，要么到头要么是相遇点，返回p即可

time complexity: O(m+n)
space complexity: O(1)
*/
func getIntersectionNodeWithoutLen(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	var p, q = headA, headB
	for p != q {
		p = nextStep(p, headB)
		q = nextStep(q, headA)
	}
	return p
}

func nextStep(p *ListNode, new *ListNode) *ListNode {
	if p == nil {
		return new
	}
	return p.Next
}