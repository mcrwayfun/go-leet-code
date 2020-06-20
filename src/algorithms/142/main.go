package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return nil
	}

	first,second := head,head
	for {
		if first.Next == nil || first.Next.Next == nil{
			return nil
		}

		first = first.Next.Next
		second = second.Next
		if first == second{
			break
		}
	}

	entrance := second
	cur := head
	for cur != entrance{
		cur = cur.Next
		entrance = entrance.Next
	}

	return cur
}

type ListNode struct {
	Val  int
	Next *ListNode
}
