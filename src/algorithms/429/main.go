package main

type Node struct {
	Val      int
	Children []*Node
}

/// Time Complexity: O(n)
/// Space Complexity: O(h)
func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	var s [][]int
	bfs(root, &s, 0)
	return s
}

func bfs(root *Node, s *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*s) == level {
		*s = append(*s, []int{})
	}

	(*s)[level] = append((*s)[level], root.Val)
	for _, v := range root.Children {
		bfs(v, s, level+1)
	}
	return
}
