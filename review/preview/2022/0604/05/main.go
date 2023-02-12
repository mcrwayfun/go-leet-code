package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue = []*TreeNode{root}
	var s *TreeNode
	var depth int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			s, queue = queue[0], queue[1:]
			if s.Left != nil {
				queue = append(queue, s.Left)
			}
			if s.Right != nil {
				queue = append(queue, s.Right)
			}
		}
		depth++
	}
	return depth
}

func maxDepthRecursion(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
