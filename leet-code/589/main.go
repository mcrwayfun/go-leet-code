package main

type Node struct {
	Val   int
	Children []*Node
}

/// Time Complexity: O(n)
/// Space Complexity: O(h)
func preorder(root *Node) []int {
	ret := make([]int, 0)
	if root == nil{
		return ret
	}

	ret = append(ret, root.Val)
	for _, v := range root.Children{
		ret = append(ret, preorder(v)...)
	}
	return ret
}