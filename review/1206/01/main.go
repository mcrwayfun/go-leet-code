package main

// time complexity: O(m+n)
// space complexity: O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	var dummy = &ListNode{}
	var prev = dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			prev.Next = list1
			list1 = list1.Next
		} else {
			prev.Next = list2
			list2 = list2.Next
		}
		prev = prev.Next
	}

	if list1 != nil {
		prev.Next = list1
		prev = prev.Next
	}
	if list2 != nil {
		prev.Next = list2
		prev = prev.Next
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}