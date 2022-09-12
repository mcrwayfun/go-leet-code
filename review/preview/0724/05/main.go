package main

// time complexity: O(n)
// space complexity: O(n)
func isBalanced(root *TreeNode) bool{
	return getHeightAndCheck(root) != -1
}

func getHeightAndCheck(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := getHeightAndCheck(root.Left)
	if left == -1 {
		return -1
	}

	right := getHeightAndCheck(root.Right)
	if right == -1 {
		return -1
	}

	if abs(left - right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
