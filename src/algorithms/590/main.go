package main

type Node struct {
	Val   int
	Children []*Node
}

/// Time Complexity: O(n)
/// Space Complexity: O(h)
func postorder(root *Node) []int {
	ret := make([]int, 0)

	if root == nil{
		return ret
	}

	for _, v := range root.Children{
		ret = append(ret, postorder(v)...)
	}
	ret = append(ret, root.Val)
	return ret
}
