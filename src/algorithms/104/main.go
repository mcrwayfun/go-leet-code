package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := 1 + maxDepth(root.Left)
	rightDepth := 1 + maxDepth(root.Right)

	if leftDepth > rightDepth {
		return leftDepth
	}

	return rightDepth
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
