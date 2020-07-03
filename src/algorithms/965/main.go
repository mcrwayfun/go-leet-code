package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/// Time Complexity: O(n)
/// Space Complexity: O(h)
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs(root, root.Val)
}

func dfs(root *TreeNode, v int) bool {
	if root == nil {
		return true
	}

	return root.Val == v && dfs(root.Left, v) && dfs(root.Right, v)
}